package gzmq

import (
	"time"

	zmq "github.com/alecthomas/gozmq"
)

type Context interface {
	NewSocket(t zmq.SocketType) (Socket, error)
	Close() error
	SetLinger(time.Duration) error
}

type context struct {
	base   zmq.Context
	linger time.Duration
	socks  map[*socket]bool
}

type Socket interface {
	zmq.Socket
}

type socket struct {
	zmq.Socket
}

// NewContext returns a new context or nil.
func NewContext() (Context, error) {
	base, err := zmq.NewContext()
	if err != nil {
		return nil, err
	}
	return &context{
		base:  base,
		socks: make(map[*socket]bool),
	}, nil
}

func (gctx *context) Close() error {
	if gctx == nil {
		return ContextIsNil
	}
	var (
		err    error
		linger = int(gctx.linger / time.Millisecond)
	)
	println("setting linger to", linger, "ms")
	for sock := range gctx.socks {
		sock_err := sock.SetSockOptInt(zmq.LINGER, linger)
		if err == nil && sock_err != nil {
		println("sock_err", sock_err.Error())
			err = sock_err
		}
	}
	for sock := range gctx.socks {
		go func() {
			start := time.Now()
			sock_err := sock.Close()
			println("closed in", time.Now().Sub(start),"ns")
			if err == nil && sock_err != nil {
		println("sock_err (B)", sock_err.Error())
				err = sock_err
			}
		}()
	}
	gctx.base.Close()
	return err
}

func (gctx *context) SetLinger(linger time.Duration) error {
	if gctx == nil {
		return ContextIsNil
	}
	gctx.linger = linger
	println("set linger to", linger, "ns")
	return nil
}

func (gctx *context) NewSocket(t zmq.SocketType) (Socket, error) {
	base, err := gctx.base.NewSocket(t)
	if err != nil {
		return nil, err
	}
	sock := &socket{
		zmq.Socket: base,
	}
	gctx.socks[sock] = true
	return sock, nil
}

type Error int

const (
	_                  = iota
	ContextIsNil Error = iota
	SocketIsNil
)

const (
	Forever time.Duration = -1 * time.Millisecond
)

func (e Error) Error() string {
	switch e {
	case ContextIsNil:
		return "gzmq: Context is nil"
	case SocketIsNil:
		return "gzmq: Socket is nil"
	}
	return "gzmq: unknown error"
}
