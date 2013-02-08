package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestContext_SetLinger(t *testing.T) {
	ctx, err := NewContext()
	if err != nil {
		t.Fatalf(err.Error())
	}
	sock, err := ctx.NewSocket(zmq.REQ)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = sock.Bind("inproc://test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	time.Sleep(5 * time.Millisecond) // make sure message reaches queue
	done := make(chan int)
	go func() {
		ctx_err := ctx.Close()
		if ctx_err != nil {
			t.Fatalf(ctx_err.Error())
		}
		done <- 1
	}()
	select {
	case <-done:
	case <-time.After(time.Millisecond):
		t.Fatalf("taking too long to close, sockets were left open?")
	}
}
