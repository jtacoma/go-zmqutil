// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzmq

import (
	"log"
	"os"
	"sync"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

type Context struct {
	base   zmq.Context
	linger time.Duration
	socks  map[*Socket]bool
	logger *log.Logger
}

type Socket struct {
	base zmq.Socket
	ctx  *Context
}

// NewContext returns a new context or nil.
func NewContext() (*Context, error) {
	base, err := zmq.NewContext()
	if err != nil {
		return nil, err
	}
	return &Context{
		base:   base,
		socks:  make(map[*Socket]bool),
		linger: -1,
	}, nil
}

// Close closes the context, blocking until the job is done.
//
// This will also propate the context's LINGER option to all sockets and close
// each of them.
func (gctx *Context) Close() error {
	if gctx == nil {
		return ContextIsNil
	}
	var (
		err error
	)
	gctx.logf("closing context: adjusting linger on sockets...")
	for sock := range gctx.socks {
		current, _ := sock.GetLinger()
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
func (gctx *Context) SetLinger(linger time.Duration) error {
	if gctx == nil {
		return ContextIsNil
	}
	gctx.linger = linger
	return nil
}

// SetVerbose enables (or disables) logging to os.Stdout.
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
func (gctx *Context) NewSocket(t zmq.SocketType) (*Socket, error) {
	base, err := gctx.base.NewSocket(t)
	if err != nil {
		gctx.logf("error while creating socket: %s", err.Error())
		return nil, err
	}
	sock := &Socket{
		base: base,
		ctx:  gctx,
	}
	gctx.socks[sock] = true
	gctx.logf("created socket %p.", sock)
	return sock, nil
}

func (gctx *Context) logf(s string, args ...interface{}) {
	if gctx.logger != nil {
		gctx.logger.Printf("[gzmq] "+s, args...)
	}
}

func (s *Socket) Close() error {
	return s.base.Close()
}

func (s *Socket) Bind(addr string) error    { return s.base.Bind(addr) }
func (s *Socket) Connect(addr string) error { return s.base.Connect(addr) }

func (s *Socket) Recv(flags zmq.SendRecvOption) ([]byte, error) {
	return s.base.Recv(flags)
}
func (s *Socket) RecvMultipart(flags zmq.SendRecvOption) ([][]byte, error) {
	return s.base.RecvMultipart(flags)
}
func (s *Socket) Send(frame []byte, flags zmq.SendRecvOption) error {
	return s.base.Send(frame, flags)
}
func (s *Socket) SendMultipart(msg [][]byte, flags zmq.SendRecvOption) error {
	return s.base.SendMultipart(msg, flags)
}

func (s *Socket) GetLinger() (time.Duration, error) {
	if s == nil {
		return -1, SocketIsNil
	}
	ms, err := s.base.GetSockOptInt(zmq.LINGER)
	if err != nil {
		return -1, err
	}
	if ms < 0 {
		return -1, nil
	}
	return time.Duration(ms) * time.Millisecond, nil
}

func (s *Socket) SetLinger(linger time.Duration) error {
	if s == nil {
		return SocketIsNil
	}
	var ms int
	if linger < 0 {
		ms = -1
	} else {
		ms = int(linger / time.Millisecond)
	}
	return s.base.SetSockOptInt(zmq.LINGER, ms)
}

type _err int

const (
	_                 = iota
	ContextIsNil _err = _err(iota)
	SocketIsNil
)

func (e _err) Error() string {
	switch e {
	case ContextIsNil:
		return "gctx: nil Context"
	case SocketIsNil:
		return "gctx: nil Socket"
	}
	return "unknown error"
}
