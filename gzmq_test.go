package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestSocket(t *testing.T) {
	var (
		context zmq.Context
		pull    zmq.Socket
		push    zmq.Socket
		polling Polling
		cpull   <-chan [][]byte
		err     error
	)
	if context, err = zmq.NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	if pull, err = context.NewSocket(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = context.NewSocket(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	polling, err = NewPolling(context)
	println("I: pushing out...")
	push.Send([]byte("test"), 0)
	println("I: pushed out")
	cpull, err = polling.Include(pull)
	select {
	case <-cpull:
		t.Logf("received message.")
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
	polling.Close()
	push.Close()
	pull.Close()
	println("I: closing context.")
	context.Close()
}
