// Copyright 2012-2013 Joshua Tacoma
//
// This file is part of ZeroMQ Utilities.
//
// ZeroMQ Utilities is free software: you can redistribute it and/or
// modify it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// ZeroMQ Utilities is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public
// License along with ZeroMQ Utilities.  If not, see
// <http://www.gnu.org/licenses/>.

package zmqutil

import (
	"log"
	"os"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

// A Context corresponds to a ØMQ context.
//
// A Context is essentially a socket factory that can be closed only after all
// the sockets it has created are also closed.
//
type Context struct {
	base   *zmq.Context
	linger time.Duration
	socks  map[*Socket]bool
	logger *log.Logger
}

// NewContext returns a new context or panics.
//
func NewContext() *Context {
	base, err := zmq.NewContext()
	if err != nil {
		panic(err.Error())
	}
	return &Context{
		base:   base,
		socks:  make(map[*Socket]bool),
		linger: -1,
	}
}

// Close closes the context, blocking until the job is done.
//
// This also propagates the context's LINGER option to all sockets that would
// linger longer and then closes each of them.
//
func (c *Context) Close() error {
	var (
		err error
	)
	c.logf("closing context: adjusting linger on sockets...")
	for sock := range c.socks {
		current, _ := sock.Linger()
		if current >= 0 && current < c.linger {
			continue
		}
		sock_err := sock.SetLinger(c.linger)
		if sock_err == zmq.ETERM {
			// "The ZeroMQ context associated with the specified socket was
			// terminated."  The context has already been closed.
			return nil
		} else if sock_err == zmq.ENOTSOCK {
			// ENOTSOCK is the error returned when the socket has already been
			// closed.  Good!
			continue
		} else if sock_err != nil {
			c.logf("closing context: error while setting linger on socket %p: %s", sock, sock_err.Error())
			if err == nil {
				err = sock_err
			}
		}
	}
	// Close each socket in a separate goroutine, effectively starting the
	// countdown of each socket's LINGER interval:
	for sock := range c.socks {
		go func(s *Socket) {
			// ZeroMQ documentation says ENOTSOCK is the only error returned
			// when closing a socket.  This is also the error returned when
			// the socket has already been closed.  Therefore, we ignore
			// errors from s.Close().
			_ = s.Close()
			c.logf("closing context: closed socket %p.", s)
		}(sock)
	}
	// As always, ZeroMQ will not return from closing a context until all
	// sockets have been closed:
	c.logf("closing context: waiting...")
	c.base.Close()
	c.logf("closing context: closed context.")
	return err
}

// SetLinger adjusts the amount of time that Close() will wait for queued
// messages to be sent.  The default is to wait forever.
//
func (c *Context) SetLinger(linger time.Duration) error {
	c.linger = linger
	return nil
}

// SetLogger sets the logger that will be used for trace logging.
//
func (c *Context) SetLogger(logger *log.Logger) {
	c.logger = logger
}

// SetVerbose enables (or disables) logging to os.Stdout.
//
// When verbose is true and a logger has already been set through SetLogger,
// this will have no effect.
//
func (c *Context) SetVerbose(verbose bool) error {
	if verbose == (c.logger != nil) {
		return nil
	}
	c.logf("verbose = %t", verbose)
	if verbose && c.logger == nil {
		c.SetLogger(log.New(os.Stdout, "", log.Lmicroseconds))
	} else if !verbose {
		c.logger = nil
	}
	c.logf("verbose = %t", c.logger != nil)
	return nil
}

// NewSocket creates a new socket and registers it to be closed when the context
// is closed.
//
// NewSocket will panic if the specified socket type is not valid, if the
// context is nil, or if there is not enough memory.
//
func (c *Context) NewSocket(t zmq.SocketType) *Socket {
	base, err := c.base.NewSocket(t)
	if err != nil {
		panic(err)
	}
	sock := &Socket{base, base}
	c.socks[sock] = true
	c.logf("created socket %p.", sock)
	return sock
}

func (c *Context) logf(s string, args ...interface{}) {
	if c.logger != nil {
		c.logger.Printf("[zmqutil] "+s, args...)
	}
}
