package main

import (
	"os"
	"time"

	zmq "github.com/alecthomas/gozmq"
	"github.com/jtacoma/go-zmqutil"
)

func main() {
	context := zmqutil.NewContext()
	defer context.Close()
	context.SetLinger(-1)
	context.SetVerbose(true)
	socket := context.NewSocket(zmq.PUSH)
	socket.MustConnect("tcp://localhost:5555")
	time.Sleep(1 * time.Second)
	println("sending:", os.Args[1])
	e := socket.SendMultipart([][]byte{[]byte(os.Args[1])}, 0)
	if e != nil {
		println("error:", e.Error())
	}
}
