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
