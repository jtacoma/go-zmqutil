// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestContext_SetLinger(t *testing.T) {
	var cases = []time.Duration{
		1 * time.Millisecond,
		4 * time.Millisecond,
		16 * time.Millisecond,
		64 * time.Millisecond,
	}
	for _, linger := range cases {
		ctx, err := NewContext()
		if err != nil {
			t.Fatalf(err.Error())
		}
		err = ctx.SetLinger(linger)
		if err != nil {
			t.Fatalf(err.Error())
		}
		sock, err := ctx.NewSocket(zmq.PUSH)
		if err != nil {
			t.Fatalf(err.Error())
		}
		err = sock.Connect("tcp://127.0.0.1:5555")
		if err != nil {
			t.Fatalf(err.Error())
		}
		go sock.Send([]byte("message1"), zmq.DONTWAIT)
		time.Sleep(time.Millisecond)
		go sock.Send([]byte("message2"), zmq.DONTWAIT)
		closing := time.Now()
		err = ctx.Close()
		if err != nil {
			t.Fatalf(err.Error())
		}
		closed := time.Now().Sub(closing)
		if closed < linger {
			t.Fatalf("closed in %s, expected >= %s", closed, linger)
		} else if linger*2 < closed {
			t.Fatalf("closed in %s, expected close to %s", closed, linger)
		}
	}
}
