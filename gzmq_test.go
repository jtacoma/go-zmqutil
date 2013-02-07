package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestContext_SetLinger(t *testing.T) {
	ctx, _ := NewContext()
	ctx.SetLinger(50 * time.Millisecond)
	sock, _ := ctx.NewSocket(zmq.PUSH)
	sock.Bind("inproc://test")
	go sock.Send([]byte("stays-in-queue"), 0)
	time.Sleep(5 * time.Millisecond) // make sure message reaches queue
	done := make(chan int)
	go func() {
		ctx.Close()
		done <- 1
	}()
	select {
	case <-done:
		t.Fatalf("closed too quickly, should have lingered.")
	case <-time.After(25 * time.Millisecond):
	}
	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("taking far too long to close when linger is set to 50ms.")
	}
}
