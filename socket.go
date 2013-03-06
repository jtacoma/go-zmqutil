// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zmqutil

import (
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

func (s *Socket) MustBind(addr string) {
	if err := s.Bind(addr); err != nil {
		panic(err.Error())
	}
}
func (s *Socket) MustConnect(addr string) {
	if err := s.Connect(addr); err != nil {
		panic(err.Error())
	}
}

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
