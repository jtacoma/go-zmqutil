package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestBridge_RegisterIn(t *testing.T) {
	unit, err := NewBridge(nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	pull, err := unit.Context().NewSocket(zmq.PULL)
	if err != nil {
		t.Fatalf(err.Error())
	}
	push, err := unit.Context().NewSocket(zmq.PUSH)
	if err != nil {
		t.Fatalf(err.Error())
	}
	pulling, err := unit.RegisterIn(pull)
	if err != nil {
		t.Fatalf(err.Error())
	}
	push.Send([]byte("test"), 0)
	select {
	case <-pulling:
		t.Logf("received message.")
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
}

func TestBridge_RegisterOut(t *testing.T) {
	t.Errorf("TODO: test this.")
}

func TestBridge_RegisterInOut(t *testing.T) {
	t.Errorf("TODO: test this.")
}

func TestBridge_Start(t *testing.T) {
	t.Errorf("TODO: test this.")
}

func TestBridge_Stop(t *testing.T) {
	t.Errorf("TODO: test this.")
}
