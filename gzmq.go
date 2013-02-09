// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzmq

import (
	"log"
	"os"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

type Context interface {
	NewSocket(t zmq.SocketType) (Socket, error)
	Close() error
	SetLinger(time.Duration) error
	SetVerbose(bool) error
}

type context struct {
	base   zmq.Context
	linger time.Duration
	socks  map[*socket]bool
	logger *log.Logger
}

type Socket interface {
	zmq.Socket
}

type socket struct {
	zmq.Socket
}

// NewContext returns a new context or nil.
func NewContext() (Context, error) {
	base, err := zmq.NewContext()
	if err != nil {
		return nil, err
	}
	return &context{
		base:  base,
		socks: make(map[*socket]bool),
	}, nil
}

func (gctx *context) Close() error {
	if gctx == nil {
		return ContextIsNil
	}
	var (
		err    error
		linger = int(gctx.linger / time.Millisecond)
	)
	gctx.logf("closing context: setting linger on sockets...")
	for sock := range gctx.socks {
		if sock_err := sock.SetSockOptInt(zmq.LINGER, linger); sock_err != nil {
			gctx.logf("closing context: error while setting linger on socket: %s", sock_err.Error())
			if err == nil {
				err = sock_err
			}
		}
	}
	for sock := range gctx.socks {
		go func() {
			if sock_err := sock.Close(); sock_err != nil && sock_err != zmq.ENOTSOCK {
				gctx.logf("closing context: error while closing socket: %s", sock_err.Error())
				if err == nil {
					err = sock_err
				}
			} else {
				gctx.logf("closing context: closed socket.")
			}
		}()
	}
	gctx.logf("closing context: closing context...")
	gctx.base.Close()
	gctx.logf("closing context: closed context.")
	return err
}

func (gctx *context) SetLinger(linger time.Duration) error {
	if gctx == nil {
		return ContextIsNil
	}
	gctx.linger = linger
	return nil
}

func (gctx *context) SetVerbose(verbose bool) error {
	if gctx == nil {
		return ContextIsNil
	}
	if verbose == (gctx.logger != nil) {
		return nil
	}
	gctx.logf("verbose = %t", gctx.logger != nil)
	if verbose && gctx.logger == nil {
		gctx.logger = log.New(os.Stdout, "", log.Lmicroseconds)
	} else if !verbose {
		gctx.logger = nil
	}
	gctx.logf("verbose = %t", gctx.logger != nil)
	return nil
}

func (gctx *context) NewSocket(t zmq.SocketType) (Socket, error) {
	base, err := gctx.base.NewSocket(t)
	if err != nil {
		gctx.logf("error while creating socket: %s", err.Error())
		return nil, err
	}
	sock := &socket{
		zmq.Socket: base,
	}
	gctx.socks[sock] = true
	gctx.logf("created socket.")
	return sock, nil
}

func (gctx *context) logf(s string, args ...interface{}) {
	if gctx.logger != nil {
		if s[len(s)-1] != '\n' {
			s += "\n"
		}
		gctx.logger.Printf("gzmq: "+s, args...)
	}
}

type Error int

const (
	_                  = iota
	ContextIsNil Error = iota
	SocketIsNil
)

func (e Error) Error() string {
	switch e {
	case ContextIsNil:
		return "gctx: nil Context"
	case SocketIsNil:
		return "gctx: nil Socket"
	}
	return "unknown error"
}
