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
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

type lingerCase struct {
	context time.Duration
	socket  time.Duration
}

func TestContext_SetLinger(t *testing.T) {
	var cases = []lingerCase{
		{1 * time.Millisecond, 1 * time.Second},
		{4 * time.Millisecond, 1 * time.Second},
		{16 * time.Millisecond, 1 * time.Second},
		{64 * time.Millisecond, 1 * time.Second},
		{1 * time.Second, 0},
		{1 * time.Second, 64 * time.Millisecond},
		{-1, 0},
		{0, -1},
	}
	for _, test := range cases {
		var linger = test.context
		if linger == -1 || linger > test.socket {
			linger = test.socket
		}
		ctx := NewContext()
		defer ctx.Close()
		err := ctx.SetLinger(test.context)
		if err != nil {
			t.Fatalf(err.Error())
		}
		sock := ctx.NewSocket(zmq.PUSH)
		sock.SetLinger(test.socket)
		err = sock.Connect("tcp://127.0.0.1:5555")
		if err != nil {
			t.Fatalf(err.Error())
		}
		go func() {
			sock.Send([]byte("message1"), 0)
			sock.Send([]byte("message2"), 0)
		}()
		time.Sleep(time.Millisecond)
		closing := time.Now()
		err = ctx.Close()
		if err != nil {
			t.Fatalf(err.Error())
		}
		closed := time.Now().Sub(closing)
		if closed < linger {
			t.Fatalf("closed in %s, expected >= %s", closed, linger)
		} else if linger*2+time.Millisecond < closed {
			t.Fatalf("closed in %s, expected close to %s", closed, linger)
		}
	}
}

func TestContext_Close(t *testing.T) {
	ctx := NewContext()
	defer ctx.Close()
	err := ctx.SetLinger(0)
	if err != nil {
		t.Fatalf(err.Error())
	}
	sock := ctx.NewSocket(zmq.PUSH)
	err = sock.Connect("tcp://127.0.0.1:5555")
	if err != nil {
		t.Fatalf(err.Error())
	}
	sock.Close()
	time.Sleep(time.Millisecond)
	err = ctx.Close()
	if err != nil {
		t.Fatalf("closing a context with an already-closed socket: %s", err)
	}
	err = ctx.Close()
	if err != nil {
		t.Fatalf("when closing an already-closed context: %s", err)
	}
}
