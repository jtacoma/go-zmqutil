// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zmqutil

import (
	"log"
	"os"
	"sync"
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
// This will also propate the context's LINGER option to all sockets and close
// each of them.
//
func (gctx *Context) Close() error {
	if gctx == nil {
		return ContextIsNil
	}
	var (
		err error
	)
	gctx.logf("closing context: adjusting linger on sockets...")
	for sock := range gctx.socks {
		current, _ := sock.Linger()
		if current >= 0 && current < gctx.linger {
			continue
		}
		if sock_err := sock.SetLinger(gctx.linger); sock_err != nil {
			gctx.logf("closing context: error while setting linger on socket %p: %s", sock, sock_err.Error())
			if err == nil {
				err = sock_err
			}
		}
	}
	var wg sync.WaitGroup
	for sock := range gctx.socks {
		wg.Add(1)
		go func(s *Socket) {
			if sock_err := s.Close(); sock_err != nil && sock_err != zmq.ENOTSOCK {
				gctx.logf("closing context: error while closing socket %p: %s", s, sock_err.Error())
				if err == nil {
					err = sock_err
				}
			} else {
				gctx.logf("closing context: closed socket %p.", s)
			}
			wg.Done()
		}(sock)
	}
	gctx.logf("closing context: waiting for sockets to close...")
	wg.Wait()
	gctx.logf("closing context: sockets closed, closing context...")
	gctx.base.Close()
	gctx.logf("closing context: closed context.")
	return err
}

// SetLinger adjusts the amount of time that Close() will wait for queued
// messages to be sent.  The default is to wait forever.
//
func (gctx *Context) SetLinger(linger time.Duration) error {
	if gctx == nil {
		return ContextIsNil
	}
	gctx.linger = linger
	return nil
}

// SetVerbose enables (or disables) logging to os.Stdout.
//
func (gctx *Context) SetVerbose(verbose bool) error {
	if gctx == nil {
		return ContextIsNil
	}
	if verbose == (gctx.logger != nil) {
		return nil
	}
	gctx.logf("verbose = %t", verbose)
	if verbose && gctx.logger == nil {
		gctx.logger = log.New(os.Stdout, "", log.Lmicroseconds)
	} else if !verbose {
		gctx.logger = nil
	}
	gctx.logf("verbose = %t", gctx.logger != nil)
	return nil
}

// NewSocket creates a new socket and registers it to be closed when the context
// is closed.
//
// NewSocket will panic if the specified socket type is not valid, if the
// context is nil, or if there is not enough memory.
//
func (gctx *Context) NewSocket(t zmq.SocketType) *Socket {
	base, err := gctx.base.NewSocket(t)
	if err != nil {
		panic(err)
	}
	sock := &Socket{
		base: base,
		ctx:  gctx,
	}
	gctx.socks[sock] = true
	gctx.logf("created socket %p.", sock)
	return sock
}

func (gctx *Context) logf(s string, args ...interface{}) {
	if gctx.logger != nil {
		gctx.logger.Printf("[zmqutil] "+s, args...)
	}
}
