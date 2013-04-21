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
	"fmt"
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func closeTestCtx(t *testing.T, ctx *Context) {
	done := make(chan int)
	go func() {
		ctx.Close()
		done <- 1
	}()
	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatalf("context is taking too long to close.")
	}
}

func ExamplePoller() {
	context := NewContext()
	defer context.Close()

	poller := NewPoller(context)

	push := context.NewSocket(zmq.PUSH)
	pull := context.NewSocket(zmq.PULL)
	push.MustBind("tcp://127.0.0.1:5555")
	pull.MustConnect("tcp://127.0.0.1:5555")

	recv := make(chan string, 1)

	poller.HandleIn(pull, func(m [][]byte) { recv <- string(m[0]) })

	push.Send([]byte("Hello!"), 0)

	poller.Poll(-1)

	fmt.Println(<-recv)

	// Output:
	// Hello!
}

func TestPoller_Poll(t *testing.T) {
	var (
		context *Context
		pull    *Socket
		push    *Socket
		poller  *Poller
		cpull   chan [][]byte
		err     error
	)
	context = NewContext()
	defer closeTestCtx(t, context)
	//context.SetVerbose(true)
	context.SetLinger(100 * time.Millisecond)
	pull = context.NewSocket(zmq.PULL)
	if err = pull.Bind("tcp://127.0.0.1:5555"); err != nil {
		t.Fatalf(err.Error())
	}
	push = context.NewSocket(zmq.PUSH)
	if err = push.Connect("tcp://127.0.0.1:5555"); err != nil {
		t.Fatalf(err.Error())
	}
	poller = NewPoller(context)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//poller.SetVerbose(true)
	push.Send([]byte("test"), 0)
	cpull = make(chan [][]byte, 2)
	poller.HandleIn(pull, func(m [][]byte) { cpull <- m })
	poller.Poll(-1)
	select {
	case <-cpull:
	default:
	}
}
