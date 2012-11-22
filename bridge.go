package gzmq

import (
	"strconv"

	zmq "github.com/alecthomas/gozmq"
)

type Chan interface {
	Bind(address string) error
	Connect(address string) error
	Close() error
	// TODO: methods to set zmq socket options
}

type In interface {
	Chan
	In() <-chan [][]byte
}

type Out interface {
	Chan
	Out() chan<- [][]byte
}

type InOut interface {
	Chan
	In() <-chan [][]byte
	Out() chan<- [][]byte
}

type Context interface {
	NewIn(s zmq.SocketType) (In, error)
	NewOut(s zmq.SocketType) (Out, error)
	NewInOut(s zmq.SocketType) (InOut, error)
	Pump() error
	Close() error
}

func NewContext() (Context, error) {
	ctx, err := zmq.NewContext()
	if err != nil {
		return nil, err
	}
	return &context{
		ctx:    ctx,
		chans:  []Chan{},
		events: make(chan event, 7),
	}, nil
}

type context struct {
	ctx    zmq.Context
	chans  []Chan
	events chan event
}

func (c *context) NewIn(t zmq.SocketType) (In, error) {
	socket, err := c.ctx.NewSocket(t)
	if err != nil {
		return nil, err
	}
	enqueue, err := c.ctx.NewSocket(zmq.PUSH)
	if err != nil {
		return nil, err
	}
	dequeue, err := c.ctx.NewSocket(zmq.PULL)
	if err != nil {
		return nil, err
	}
	addr := newInprocAddress()
	if err := enqueue.Bind(addr); err != nil {
		return nil, err
	}
	if err := dequeue.Connect(addr); err != nil {
		return nil, err
	}
	result := &chanIn{
		chanBase{
			socket,
			make(chan event, 7),
		},
		enqueue, dequeue,
		make(chan [][]byte, 7),
	}
	c.chans = append(c.chans, result)
	return result, nil
}

func (c *context) NewOut(t zmq.SocketType) (Out, error) {
	socket, err := c.ctx.NewSocket(t)
	if err != nil {
		return nil, err
	}
	result := &chanOut{
		chanBase{
			socket,
			make(chan event, 7),
		},
		make(chan [][]byte, 7),
	}
	c.chans = append(c.chans, result)
	return result, nil
}

func (c *context) NewInOut(t zmq.SocketType) (InOut, error) {
	socket, err := c.ctx.NewSocket(t)
	if err != nil {
		return nil, err
	}
	enqueue, err := c.ctx.NewSocket(zmq.PUSH)
	if err != nil {
		return nil, err
	}
	dequeue, err := c.ctx.NewSocket(zmq.PULL)
	if err != nil {
		return nil, err
	}
	addr := newInprocAddress()
	if err := enqueue.Bind(addr); err != nil {
		return nil, err
	}
	if err := dequeue.Connect(addr); err != nil {
		return nil, err
	}
	return &chanInOut{
		chanBase{
			socket,
			make(chan event, 7),
		},
		enqueue, dequeue,
		make(chan [][]byte, 7),
		make(chan [][]byte, 7),
	}, nil
}

func (c *context) Pump() error {
	println("I: Pump() [start]")
	for _, ch := range c.chans {
		switch ch.(type) {
		case *chanOut:
			go ch.(*chanOut).pumpOut()
		case *chanInOut:
			go ch.(*chanInOut).pumpOut()
		}
		switch ch.(type) {
		case *chanIn:
			go ch.(*chanIn).pumpIn()
		case *chanInOut:
			go ch.(*chanInOut).pumpIn()
		}
	}
	for e := range c.events {
		switch e {
		case closing:
			for _, ch := range c.chans {
				switch ch.(type) {
				case *chanIn:
					println("I: enqueuing close command for in")
					ch.(*chanIn).enqueue.SendMultipart([][]byte{[]byte{byte(closing)}}, 0)
				case *chanInOut:
					ch.(*chanInOut).events <- e
				case *chanOut:
					ch.(*chanOut).events <- e
				}
			}
			for _, ch := range c.chans {
				switch ch.(type) {
				case *chanIn:
					ch.(*chanIn).enqueue.Close()
				}
			}
		}
	}
	println("I: closing zmq context...")
	c.ctx.Close()
	println("I: Pump() [end]")
	return nil
}

func (c *context) Close() error {
	c.events <- closing
	c.ctx.Close()
	return nil
}

type event int

const (
	_             = iota
	closing event = iota
)

type chanBase struct {
	socket zmq.Socket
	events chan event
}

func (c *chanBase) Bind(address string) error {
	return c.socket.Bind(address)
}

func (c *chanBase) Connect(address string) error {
	return c.socket.Connect(address)
}

func (c *chanBase) Close() error {
	// TODO: this is probably not robust...?
	return c.socket.Close()
}

type chanIn struct {
	chanBase
	enqueue zmq.Socket
	dequeue zmq.Socket
	cin     chan [][]byte
}

func (c *chanIn) In() <-chan [][]byte { return c.cin }

func (c *chanIn) pumpIn() {
	items := zmq.PollItems{
		zmq.PollItem{
			Socket: c.chanBase.socket,
			Events: zmq.POLLIN,
		},
		zmq.PollItem{
			Socket: c.dequeue,
			Events: zmq.POLLIN,
		},
	}
	for {
		_, err := zmq.Poll(items, -1)
		if err == zmq.ETERM {
			println("I: closing (in) due to context termination")
			break
		} else if err != nil {
			panic(err.Error())
		}
		if (items[0].REvents & zmq.POLLIN) != 0 {
			println("I: in received a message")
			msg, err := c.chanBase.socket.RecvMultipart(0)
			if err != nil {
				panic(err.Error())
			}
			c.cin <- msg
		}
		if (items[1].REvents & zmq.POLLIN) != 0 {
			println("I: in received an event")
			msg, err := c.dequeue.RecvMultipart(0)
			if err != nil {
				panic(err.Error())
			}
			if len(msg[0]) == 0 {
				continue
			}
			switch event(int(msg[0][0])) {
			case closing:
				println("I: closing (in)")
				break
			}
		}
	}
	c.dequeue.Close()
	c.socket.Close()
}

type chanOut struct {
	chanBase
	cout chan [][]byte
}

func (c *chanOut) Out() chan<- [][]byte { return c.cout }

func (c *chanOut) pumpOut() {
	println("I: pumping out...")
	for {
		select {
		case msg := <-c.cout:
			println("I: sending multipart message...")
			c.socket.SendMultipart(msg, 0) // TODO: handle error
			println("I: sent multipart message")
		case e := <-c.events:
			switch e {
			case closing:
				println("I: closing (out)")
				c.socket.Close()
				break
			}
		}
		println("I: pumped...")
	}
	println("I: closed.")
}

type chanInOut struct {
	chanBase
	enqueue zmq.Socket
	dequeue zmq.Socket
	cin     chan [][]byte
	cout    chan [][]byte
}

func (c *chanInOut) In() <-chan [][]byte  { return c.cin }
func (c *chanInOut) Out() chan<- [][]byte { return c.cout }

func (c *chanInOut) pumpIn() {
	panic("not implemented")
}

func (c *chanInOut) pumpOut() {
	for {
		select {
		case msg := <-c.cout:
			c.enqueue.SendMultipart(msg, 0) // TODO: handle error
		case e := <-c.events:
			switch e {
			case closing:
				// TODO!
				break
			}
		}
	}
}

var inprocNext = 1

func newInprocAddress() string {
	inprocNext += 1
	return "inproc://github.com/jtacoma/gzmq/" + strconv.Itoa(inprocNext-1)
}
