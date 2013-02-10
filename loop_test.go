// Copyright 2012-2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzmq

import (
	"fmt"
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func closeTestCtx(t *testing.T, ctx Context) {
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
	defer closeTestCtx(nil, context)

	loop, _ := NewLoop(context)

	cli, _ := context.NewSocket(zmq.REQ)
	srv, _ := context.NewSocket(zmq.ROUTER)
	srv.Bind("tcp://127.0.0.1:5557")
	cli.Connect("tcp://127.0.0.1:5557")

	loop.HandleFunc(srv, zmq.POLLIN, func(e *SocketEvent) error {
		msg, err := srv.RecvMultipart(0)
		if err != nil {
			return err
		}
		return srv.SendMultipart(msg, 0)
	})

	recv := make(chan string, 2)
	loop.HandleFunc(cli, zmq.POLLIN, func(e *SocketEvent) error {
		msg, err := cli.Recv(0)
		if err != nil {
			return err
		}
		recv <- string(msg)
		return nil
	})

	loop.Sync(func() {
		cli.Send([]byte("Echo!"), 0)
	})

	// Receive response:
	msg := <-recv

	fmt.Println(msg)

	// Output:
	// Echo!
}

func TestLoop(t *testing.T) {
	var (
		context Context
		pull    Socket
		push    Socket
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
	loop, err = NewLoop(context)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//loop.SetVerbose(true)
	push.Send([]byte("test"), 0)
	cpull = make(chan [][]byte)
	loop.HandleFunc(pull, zmq.POLLIN, func(e *SocketEvent) error {
		msg, err := pull.RecvMultipart(0)
		if err != nil {
			return err
		}
		cpull <- msg
		return nil
	})
	select {
	case <-cpull:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
}

func TestLoop_Sync(t *testing.T) {
	var (
		context    Context
		reQ, reP   Socket
		loop       *Loop
		creQ, creP chan [][]byte
		err        error
	)
	if context, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	//context.SetVerbose(true)
	if reQ, err = context.NewSocket(zmq.REQ); err != nil {
		t.Fatalf(err.Error())
	}
	if err = reQ.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if reP, err = context.NewSocket(zmq.REP); err != nil {
		t.Fatalf(err.Error())
	}
	if err = reP.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	loop, err = NewLoop(context)
	//loop.SetVerbose(true)
	creP = make(chan [][]byte, 2)
	repHandler := loop.HandleFunc(reP, zmq.POLLIN, func(e *SocketEvent) error {
		if msg, err := e.Socket.RecvMultipart(0); err != nil {
			return err
		} else {
			creP <- msg
		}
		return nil
	})
	creQ = make(chan [][]byte, 2)
	loop.HandleFunc(reQ, zmq.POLLIN, func(e *SocketEvent) error {
		if msg, err := e.Socket.RecvMultipart(0); err != nil {
			return err
		} else {
			creQ <- msg
		}
		return nil
	})
	loop.Sync(func() { reQ.Send([]byte("request"), 0) })
	done := false
	for !done {
		select {
		case <-creP:
			loop.Sync(func() { reP.Send([]byte("response"), 0) })
			err = loop.HandleEnd(reP, zmq.POLLIN, repHandler)
			if err != nil {
				t.Fatalf(err.Error())
			}
		case <-creQ:
			done = true
		case <-time.After(10 * time.Millisecond):
			t.Errorf("failed to receive message within timeout.")
			done = true
		}
	}
}
