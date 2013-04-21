package main

import (
	"log"
	"os"
	"time"

	zmq "github.com/alecthomas/gozmq"
	"github.com/jtacoma/go-zmqutil"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	context := zmqutil.NewContext()
	defer context.Close()
	context.SetLinger(50 * time.Second)
	context.SetLogger(logger)
	socket := context.NewSocket(zmq.PULL)
	poller := zmqutil.NewPoller(context)
	poller.HandleIn(socket, func(m [][]byte) {
		logger.Println("received:", string(m[0]))
		if string(m[0]) == "STOP" {
			poller.Unhandle(socket)
			logger.Println("stopping...")
		}
	})
	socket.MustBind("tcp://*:5555")
	poller.Run()
}
