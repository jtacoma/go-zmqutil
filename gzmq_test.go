package gzmq

import (
	"testing"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

func TestContext_SetLinger(t *testing.T) {
	var cases = []time.Duration{
		1 * time.Millisecond,
		4 * time.Millisecond,
		16 * time.Millisecond,
		64 * time.Millisecond,
	}
	for _, linger := range cases {
		ctx, _ := NewContext()
		ctx.SetLinger(linger)
		sock, _ := ctx.NewSocket(zmq.PUSH)
		sock.Connect("tcp://127.0.0.1:5555")
		go sock.Send([]byte("message1"), zmq.DONTWAIT)
		time.Sleep(time.Millisecond)
		go sock.Send([]byte("message2"), zmq.DONTWAIT)
		closing := time.Now()
		ctx.Close()
		closed := time.Now().Sub(closing)
		if closed < linger {
			t.Fatalf("closed in %s, expected >= %s", closed, linger)
		} else if linger*2 < closed {
			t.Fatalf("closed in %s, expected close to %s", closed, linger)
		}
	}
}
