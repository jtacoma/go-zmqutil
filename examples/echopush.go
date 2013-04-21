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
	context.SetLinger(-1)
	context.SetLogger(logger)
	socket := context.NewSocket(zmq.PUSH)
	socket.MustConnect("tcp://localhost:5555")
	time.Sleep(1 * time.Second)
	logger.Println("sending:", os.Args[1])
	e := socket.SendMultipart([][]byte{[]byte(os.Args[1])}, 0)
	if e != nil {
		logger.Println("error:", e.Error())
	}
}
