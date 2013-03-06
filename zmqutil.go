// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zmqutil

import zmq "github.com/alecthomas/gozmq"

type _err int

const (
	_                 = iota
	ContextIsNil _err = _err(iota)
	SocketIsNil
	NotImplemented
)

// TODO: these should be const:
var (
	PAIR   = zmq.SocketType(zmq.PAIR)
	PUB    = zmq.SocketType(zmq.PUB)
	SUB    = zmq.SocketType(zmq.SUB)
	REQ    = zmq.SocketType(zmq.REQ)
	REP    = zmq.SocketType(zmq.REP)
	DEALER = zmq.SocketType(zmq.DEALER)
	ROUTER = zmq.SocketType(zmq.ROUTER)
	PULL   = zmq.SocketType(zmq.PULL)
	PUSH   = zmq.SocketType(zmq.PUSH)
	XPUB   = zmq.SocketType(zmq.XPUB)
	XSUB   = zmq.SocketType(zmq.XSUB)
)

// TODO: these should be const:
var (
	POLLIN  = zmq.PollEvents(zmq.POLLIN)
	POLLOUT = zmq.PollEvents(zmq.POLLOUT)
	POLLERR = zmq.PollEvents(zmq.POLLERR)
)

func (e _err) Error() string {
	switch e {
	case ContextIsNil:
		return "zmqutil: nil Context"
	case SocketIsNil:
		return "zmqutil: nil Socket"
	case NotImplemented:
		return "zmqutil: this function is not yet implemented."
	}
	return "unknown error"
}
