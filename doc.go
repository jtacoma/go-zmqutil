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

A reactor loop that lets event handlers be attached to sockets.

    import (
        "errors"
        "time"

        zmq "github.com/alecthomas/gozmq"
        "github.com/jtacoma/go-zmqutil"
    )

    func main() {
        context := zmqutil.NewContext()
        defer context.Close()
        context.SetLinger(1 * time.Second)
        socket, _ := context.NewSocket(zmq.SUB)
        poller, _ := zmqutil.NewPoller(context)
        poller.HandlFunc(socket, zmq.POLLIN, func (e *zmqutil.SocketEvent) error {
            println(string(msg[0]))
            if (string(msg[0]) == "STOP") {
                e.Fault = errors.New("received 'STOP'")
                println("Stopping...")
            }
        })
        socket.Bind("tcp://localhost:5555")
        poller.Run()
    }

*/
package zmqutil
