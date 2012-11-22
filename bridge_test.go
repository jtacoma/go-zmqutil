package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestContext_PushPull(t *testing.T) {
	var (
		unit Context
		pull In
		push Out
		err  error
	)
	if unit, err = NewContext(); err != nil {
		t.Fatalf(err.Error())
	}
	if pull, err = unit.NewIn(zmq.PULL); err != nil {
		t.Fatalf(err.Error())
	}
	if err = pull.Bind("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	if push, err = unit.NewOut(zmq.PUSH); err != nil {
		t.Fatalf(err.Error())
	}
	if err = push.Connect("inproc://test"); err != nil {
		t.Fatalf(err.Error())
	}
	go func() {
		println("I: pumping context")
		if err := unit.Pump(); err != nil {
			t.Errorf(err.Error())
		}
		println("!!!")
	}()
	time.Sleep(100 * time.Millisecond)
	println("I: pushing out...")
	push.Out() <- [][]byte{[]byte("test")}
	println("I: pushed out")
	select {
	case <-pull.In():
		t.Logf("received message.")
	case <-time.After(10 * time.Millisecond):
		t.Errorf("failed to receive message within timeout.")
	}
	if err = unit.Close(); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestContext_NewOut(t *testing.T) {
	t.Errorf("TODO: test this.")
}

func TestContext_NewInOut(t *testing.T) {
	t.Errorf("TODO: test this.")
}

func TestContext_Pump(t *testing.T) {
	t.Errorf("TODO: test this.")
}

func TestContext_Close(t *testing.T) {
	t.Errorf("TODO: test this.")
}
