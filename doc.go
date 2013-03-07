// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package zmqutil implements some ØMQ (http://www.zeromq.org) abstractions and
utilities.

A context from this package remembers its sockets and has its own Linger
option.  When a context is closed, it will set the Linger option on each
socket then close them all.

All socket options are available through option-specific getter/setter methods.

An additonal type, Poller, provides reactor loop that lets event handlers be
attached to sockets.

	package main

	import (
		"errors"
		"time"

		zmq "github.com/alecthomas/gozmq"
		"github.com/jtacoma/go-zmqutil"
	)

	func echo(e *zmqutil.Event, m [][]byte) {
		println("received:", string(m[0]))
		if string(m[0]) == "STOP" {
			e.Fault = errors.New("received 'STOP'")
			println("stopping...")
		}
	}

	func main() {
		context := zmqutil.NewContext()
		defer context.Close()
		context.SetLinger(50 * time.Second)
		context.SetVerbose(true)
		socket := context.NewSocket(zmq.PULL)
		poller := zmqutil.NewPoller(context)
		poller.Handle(socket, zmq.POLLIN, zmqutil.NewMessageHandler(echo))
		socket.MustBind("tcp://*:5555")
		poller.Run()
	}

*/
package zmqutil
