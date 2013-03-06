package main

import (
	"errors"
	"time"

	zmq "github.com/alecthomas/gozmq"
	"github.com/jtacoma/go-zmqutil"
)

func main() {
	context := zmqutil.NewContext()
	defer context.Close()
	context.SetLinger(50 * time.Second)
	context.SetVerbose(true)
	socket := context.NewSocket(zmq.PULL)
	poller := zmqutil.NewPoller(context)
	poller.HandleFunc(socket, zmq.POLLIN, func(e *zmqutil.Event) {
		println("received:", string(e.Message[0]))
		if string(e.Message[0]) == "STOP" {
			e.Fault = errors.New("received 'STOP'")
			println("stopping...")
		}
	})
	socket.MustBind("tcp://*:5555")
	poller.Run()
}
