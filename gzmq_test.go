package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestSocket(t *testing.T) {
	var (
		context zmq.Context
		pull    Socket
		push    Socket
		err     error
	)
	if context, err = zmq.NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	if pull, err = NewSocket(context, zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = NewSocket(context, zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	go pull.Pump()
	go push.Pump()
	time.Sleep(100 * time.Millisecond)
	println("I: pushing out...")
	push.Send() <- [][]byte{[]byte("test")}
	println("I: pushed out")
	select {
	case <-pull.Recv():
		t.Logf("received message.")
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
	push.Close()
	pull.Close()
	println("I: closing context.")
	context.Close()
}
