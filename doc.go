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

/*
Package zmqutil implements some ØMQ (http://www.zeromq.org) abstractions and
utilities.

A context from this package remembers its sockets and has its own Linger
option.  When a context is closed, it will set the Linger option on each
socket then close them all.

All socket options are available through option-specific getter/setter methods.

An additonal type, Poller, provides a reactor loop that lets event handlers be
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
