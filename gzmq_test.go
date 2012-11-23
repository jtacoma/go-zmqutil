package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestPolling(t *testing.T) {
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
	push.Send([]byte("test"), 0)
	cpull, err = polling.Include(pull)
	select {
	case <-cpull:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
	if err = polling.Close(); err != nil {
		t.Fatalf(err.Error())
	}
	if err = push.Close(); err != nil {
		t.Fatalf(err.Error())
	}
	if err = pull.Close(); err != nil {
		t.Fatalf(err.Error())
	}
	context.Close()
}

func TestPolling_Sync(t *testing.T) {
	var (
		context    zmq.Context
		reQ, reP   zmq.Socket
		polling    Polling
		creQ, creP <-chan [][]byte
		err        error
	)
	if context, err = zmq.NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	if reQ, err = context.NewSocket(zmq.REQ); err != nil {
		t.Fatalf(err.Error())
	}
	if err = reQ.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if reP, err = context.NewSocket(zmq.REP); err != nil {
		t.Fatalf(err.Error())
	}
	if err = reP.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	polling, err = NewPolling(context)
	if creP, err = polling.Include(reP); err != nil {
		t.Fatalf(err.Error())
	}
	if creQ, err = polling.Include(reQ); err != nil {
		t.Fatalf(err.Error())
	}
	polling.Sync(func() { reQ.Send([]byte("request"), 0) })
	done := false
	for !done {
		select {
		case <-creP:
			polling.Sync(func() { reP.Send([]byte("response"), 0) })
		case <-creQ:
			done = true
		case <-time.After(10 * time.Millisecond):
			t.Errorf("failed to receive message within timeout.")
			done = true
		}
	}
	if err = polling.Close(); err != nil {
		t.Fatalf(err.Error())
	}
	if err = reP.Close(); err != nil {
		t.Fatalf(err.Error())
	}
	if err = reQ.Close(); err != nil {
		t.Fatalf(err.Error())
	}
	context.Close()
}
