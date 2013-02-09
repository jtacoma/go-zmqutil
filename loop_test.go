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

func ExampleLoop() {
	context, _ := NewContext()
	defer context.Close()

	loop, _ := NewLoop(context, 1)
	defer loop.Close()

	cli, _ := context.NewSocket(zmq.REQ)
	defer cli.Close()
	srv, _ := context.NewSocket(zmq.ROUTER)
	defer srv.Close()
	srv.Bind("inproc://example")
	cli.Connect("inproc://example")

	cli_send, _ := loop.NewSending(cli)
	srv_recv, _ := loop.Start(srv, 1)
	srv_send, _ := loop.NewSending(srv)
	cli_recv, _ := loop.Start(cli, 1)

	go func() {
		// Process requests:
		for msg := range srv_recv {
			srv_send <- msg
		}
	}()

	// Send request:
	cli_send <- [][]byte{[]byte("Echo!")}

	// Receive response:
	msg := <-cli_recv

	fmt.Println(string(msg[0]))

	// Output:
	// Echo!
}

func TestNewSending(t *testing.T) {
	var (
		context Context
		pull    Socket
		push    Socket
		loop    Loop
		cpush   chan<- [][]byte
		cpull   <-chan [][]byte
		err     error
	)
	if context, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	if pull, err = context.NewSocket(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	defer pull.Close()
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = context.NewSocket(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	defer push.Close()
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	cpush, _ = NewSending(push)
	cpush <- [][]byte{[]byte("test")}
	loop, err = NewLoop(context, 1)
	defer loop.Close()
	cpull, _ = loop.Start(pull, 1)
	select {
	case <-cpull:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
}

func TestLoop(t *testing.T) {
	var (
		context Context
		pull    Socket
		push    Socket
		loop    Loop
		cpull   <-chan [][]byte
		err     error
	)
	if context, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	if pull, err = context.NewSocket(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	defer pull.Close()
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = context.NewSocket(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	defer push.Close()
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	loop, err = NewLoop(context, 1)
	defer loop.Close()
	push.Send([]byte("test"), 0)
	cpull, err = loop.Start(pull, 1)
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
		loop       Loop
		creQ, creP <-chan [][]byte
		err        error
	)
	if context, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	if reQ, err = context.NewSocket(zmq.REQ); err != nil {
		t.Fatalf(err.Error())
	}
	defer reQ.Close()
	if err = reQ.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if reP, err = context.NewSocket(zmq.REP); err != nil {
		t.Fatalf(err.Error())
	}
	defer reP.Close()
	if err = reP.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	loop, err = NewLoop(context, 1)
	defer loop.Close()
	if creP, err = loop.Start(reP, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if creQ, err = loop.Start(reQ, 1); err != nil {
		t.Fatalf(err.Error())
	}
	loop.Sync(func() { reQ.Send([]byte("request"), 0) })
	done := false
	for !done {
		select {
		case <-creP:
			loop.Sync(func() { reP.Send([]byte("response"), 0) })
			err = loop.Stop(reP)
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

func TestLoop_Sending(t *testing.T) {
	var (
		context            Context
		reQ, reP           Socket
		loop               Loop
		creQrecv, crePrecv <-chan [][]byte
		creQsend, crePsend chan<- [][]byte
		err                error
	)
	if context, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	if reQ, err = context.NewSocket(zmq.REQ); err != nil {
		t.Fatalf(err.Error())
	}
	defer reQ.Close()
	if err = reQ.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if reP, err = context.NewSocket(zmq.REP); err != nil {
		t.Fatalf(err.Error())
	}
	defer reP.Close()
	if err = reP.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	loop, err = NewLoop(context, 1)
	defer loop.Close()
	if crePrecv, err = loop.Start(reP, 1); err != nil {
		t.Fatalf(err.Error())
	}
	if creQrecv, err = loop.Start(reQ, 1); err != nil {
		t.Fatalf(err.Error())
	}
	crePsend, _ = loop.NewSending(reP)
	creQsend, _ = loop.NewSending(reQ)
	creQsend <- [][]byte{[]byte("request")}
	done := false
	for !done {
		select {
		case <-crePrecv:
			crePsend <- [][]byte{[]byte("response")}
			err = loop.Stop(reP)
			if err != nil {
				t.Fatalf(err.Error())
			}
		case <-creQrecv:
			done = true
		case <-time.After(10 * time.Millisecond):
			t.Errorf("failed to receive message within timeout.")
			done = true
		}
	}
}
