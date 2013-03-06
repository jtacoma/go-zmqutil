// Copyright 2012-2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
	push.Bind("tcp://127.0.0.1:5557")
	pull.Connect("tcp://127.0.0.1:5557")

	recv := make(chan string, 2)

	poller.HandleFunc(pull, zmq.POLLIN, func(e *SocketEvent) {
		recv <- string(e.Message[0])
	})

	push.Send([]byte("Echo!"), 0)

	poller.Poll(1 * time.Second)

	fmt.Println(<-recv)

	// Output:
	// Echo!
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
	poller.HandleFunc(pull, zmq.POLLIN, func(e *SocketEvent) {
		cpull <- e.Message
	})
	poller.Poll(-1)
	select {
	case <-cpull:
	default:
	}
}
