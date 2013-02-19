// Copyright 2012-2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gozmqutil

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
		if t != nil {
			t.Fatalf("context is taking too long to close.")
		} else {
			println("ERROR: context is taking too long to close.")
		}
	}
}

func ExampleLoop() {
	context, _ := NewContext()
	defer context.Close()

	loop := NewLoop(context)

	cli, _ := context.NewSocket(zmq.REQ)
	srv, _ := context.NewSocket(zmq.ROUTER)
	srv.Bind("tcp://127.0.0.1:5557")
	cli.Connect("tcp://127.0.0.1:5557")

	loop.HandleFunc(srv, zmq.POLLIN, func(e *SocketEvent) {
		e.Fault = srv.SendMultipart(e.Message, 0)
	})

	recv := make(chan string, 2)
	loop.HandleFunc(cli, zmq.POLLIN, func(e *SocketEvent) {
		recv <- string(e.Message[0])
	})

	cli.Send([]byte("Echo!"), 0)

	go loop.Run()

	// Receive response:
	msg := <-recv

	fmt.Println(msg)

	// Output:
	// Echo!
}

func TestLoop_Step(t *testing.T) {
	var (
		context *Context
		pull    *Socket
		push    *Socket
		loop    *Loop
		cpull   chan [][]byte
		err     error
	)
	if context, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer closeTestCtx(t, context)
	//context.SetVerbose(true)
	context.SetLinger(100 * time.Millisecond)
	if pull, err = context.NewSocket(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	if err = pull.Bind("tcp://127.0.0.1:5555"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = context.NewSocket(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	if err = push.Connect("tcp://127.0.0.1:5555"); err != nil {
		t.Fatalf(err.Error())
	}
	loop = NewLoop(context)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//loop.SetVerbose(true)
	push.Send([]byte("test"), 0)
	cpull = make(chan [][]byte, 2)
	loop.HandleFunc(pull, zmq.POLLIN, func(e *SocketEvent) {
		cpull <- e.Message
	})
	loop.Step(-1)
	select {
	case <-cpull:
	default:
	}
}
