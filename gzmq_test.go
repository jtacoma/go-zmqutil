package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestNewSending(t *testing.T) {
	var (
		context zmq.Context
		pull    zmq.Socket
		push    zmq.Socket
		polling Polling
		cpush   chan<- [][]byte
		cpull   <-chan [][]byte
		err     error
	)
	if context, err = zmq.NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	if pull, err = context.NewSocket(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	defer pull.Close()
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = context.NewSocket(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	defer push.Close()
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	cpush, _ = NewSending(push)
	cpush <- [][]byte{[]byte("test")}
	polling, err = NewPolling(context)
	defer polling.Close()
	cpull, _ = polling.Start(pull)
	select {
	case <-cpull:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
}

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
	defer context.Close()
	if pull, err = context.NewSocket(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	defer pull.Close()
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = context.NewSocket(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	defer push.Close()
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	polling, err = NewPolling(context)
	defer polling.Close()
	push.Send([]byte("test"), 0)
	cpull, err = polling.Start(pull)
	select {
	case <-cpull:
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
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
	defer context.Close()
	if reQ, err = context.NewSocket(zmq.REQ); err != nil {
		t.Fatalf(err.Error())
	}
	defer reQ.Close()
	if err = reQ.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if reP, err = context.NewSocket(zmq.REP); err != nil {
		t.Fatalf(err.Error())
	}
	defer reP.Close()
	if err = reP.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	polling, err = NewPolling(context)
	defer polling.Close()
	if creP, err = polling.Start(reP); err != nil {
		t.Fatalf(err.Error())
	}
	if creQ, err = polling.Start(reQ); err != nil {
		t.Fatalf(err.Error())
	}
	polling.Sync(func() { reQ.Send([]byte("request"), 0) })
	done := false
	for !done {
		select {
		case <-creP:
			polling.Sync(func() { reP.Send([]byte("response"), 0) })
			err = polling.Stop(reP)
			if err != nil {
				t.Fatalf(err.Error())
			}
		case <-creQ:
			done = true
		case <-time.After(10 * time.Millisecond):
			t.Errorf("failed to receive message within timeout.")
			done = true
		}
	}
}

func TestPolling_Sending(t *testing.T) {
	var (
		context            zmq.Context
		reQ, reP           zmq.Socket
		polling            Polling
		creQrecv, crePrecv <-chan [][]byte
		creQsend, crePsend chan<- [][]byte
		err                error
	)
	if context, err = zmq.NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	defer context.Close()
	if reQ, err = context.NewSocket(zmq.REQ); err != nil {
		t.Fatalf(err.Error())
	}
	defer reQ.Close()
	if err = reQ.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if reP, err = context.NewSocket(zmq.REP); err != nil {
		t.Fatalf(err.Error())
	}
	defer reP.Close()
	if err = reP.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	polling, err = NewPolling(context)
	defer polling.Close()
	if crePrecv, err = polling.Start(reP); err != nil {
		t.Fatalf(err.Error())
	}
	if creQrecv, err = polling.Start(reQ); err != nil {
		t.Fatalf(err.Error())
	}
	crePsend, _ = polling.NewSending(reP)
	creQsend, _ = polling.NewSending(reQ)
	creQsend <- [][]byte{[]byte("request")}
	done := false
	for !done {
		select {
		case <-crePrecv:
			crePsend <- [][]byte{[]byte("response")}
			err = polling.Stop(reP)
			if err != nil {
				t.Fatalf(err.Error())
			}
		case <-creQrecv:
			done = true
		case <-time.After(10 * time.Millisecond):
			t.Errorf("failed to receive message within timeout.")
			done = true
		}
	}
}
