// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gozmqutil

import (
	"time"

	zmq "github.com/alecthomas/gozmq"
)

type Socket struct {
	base zmq.Socket
	ctx  *Context
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

func (s *Socket) Linger() (time.Duration, error) {
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
