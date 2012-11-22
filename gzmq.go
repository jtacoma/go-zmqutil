package gzmq

import (
	"strconv"

	zmq "github.com/alecthomas/gozmq"
)

type Socket interface {
	Bind(address string) error
	Connect(address string) error
	Pump()
	Recv() <-chan [][]byte
	Send() chan<- [][]byte
	Close() error
}

type socket struct {
	main    zmq.Socket    // the socket this struct abstracts
	cmdSend zmq.Socket    // inproc for sending commands to pumpSock()
	cmdRecv zmq.Socket    // inproc for receving commands in pumpSock()
	recv    chan [][]byte // messages that have been received
	send    chan [][]byte // messages waiting to be sent
	cmd     chan [][]byte // commands to be processed
}

func NewSocket(c zmq.Context, t zmq.SocketType) (Socket, error) {
	main, err := c.NewSocket(t)
	if err != nil {
		return nil, err
	}
	cmdSend, err := c.NewSocket(zmq.PUSH)
	if err != nil {
		return nil, err
	}
	cmdRecv, err := c.NewSocket(zmq.PULL)
	if err != nil {
		return nil, err
	}
	cmdAddr := newInprocAddress()
	err = cmdSend.Bind(cmdAddr)
	if err != nil {
		return nil, err
	}
	err = cmdRecv.Connect(cmdAddr)
	if err != nil {
		return nil, err
	}
	result := &socket{
		main:    main,
		cmdSend: cmdSend,
		cmdRecv: cmdRecv,
		recv:    make(chan [][]byte, 2),
		send:    make(chan [][]byte, 2),
		cmd:     make(chan [][]byte, 2),
	}
	return result, nil
}

func (s *socket) Bind(address string) error    { return s.main.Bind(address) }
func (s *socket) Connect(address string) error { return s.main.Connect(address) }
func (s *socket) Recv() <-chan [][]byte        { return s.recv }
func (s *socket) Send() chan<- [][]byte        { return s.send }

func (s *socket) Close() error {
	println("I: sending close to pumpChan()", s)
	s.cmd <- [][]byte{[]byte{byte(closing)}}
	return s.main.Close()
}

func (s *socket) Pump() {
	go s.pumpSock()
	s.pumpChan()
	println("I: closing socket", s)
	s.main.Close()
	println("I: closed socket", s)
}

// pumpChan pumps messages out from channels in to sockets.
func (s *socket) pumpChan() {
	defer s.cmdSend.Close()
	var stopping = false
	for !stopping {
		println("I: pumpChan() back in the select again", s)
		select {
		case msg := <-s.send:
			s.main.SendMultipart(msg, 0) // TODO: handle error
		case e := <-s.cmd:
			switch event(e[0][0]) {
			case closing:
				println("I: sending close to pumpSock()", s)
				s.cmdSend.SendMultipart([][]byte{[]byte{byte(closing)}}, 0)
				println("I: breaking pumpChan()", s)
				stopping = true
			}
		}
	}
	println("I: closed pumpChan()", s)
}

// pumpSock pumps messages in from sockets and out to channels.
func (s *socket) pumpSock() {
	defer s.cmdRecv.Close()
	defer s.main.Close()
	items := zmq.PollItems{
		zmq.PollItem{
			Socket: s.cmdRecv,
			Events: zmq.POLLIN,
		},
	}
	typ, err := s.main.GetSockOptUInt64(zmq.TYPE)
	if err != nil {
		panic(err.Error())
	}
	switch zmq.SocketType(typ) {
	case zmq.SUB, zmq.REQ, zmq.REP, zmq.DEALER, zmq.ROUTER, zmq.PULL:
		items = append(items, zmq.PollItem{
			Socket: s.main,
			Events: zmq.POLLIN,
		})
	}
	for {
		_, err := zmq.Poll(items, -1)
		if err == zmq.ETERM {
			println("I: closing pumpSock() [context terminated]")
			break
		} else if err != nil {
			panic(err.Error())
		}
		if (items[0].REvents & zmq.POLLIN) != 0 {
			println("I: in received an event")
			msg, err := s.cmdRecv.RecvMultipart(0)
			if err != nil {
				panic(err.Error())
			}
			if len(msg[0]) == 0 {
				continue
			}
			switch event(int(msg[0][0])) {
			case closing:
				println("I: closing pumpSock()")
				break
			}
		}
		if len(items) > 1 && (items[1].REvents&zmq.POLLIN) != 0 {
			println("I: in received a message")
			msg, err := s.main.RecvMultipart(0)
			if err != nil {
				panic(err.Error())
			}
			s.recv <- msg
		}
	}
	println("I: closed pumpSock()", s)
}

type event int

const (
	_             = iota
	closing event = iota
)

var inprocNext = 1

func newInprocAddress() string {
	inprocNext += 1
	return "inproc://github.com/jtacoma/gzmq/" + strconv.Itoa(inprocNext-1)
}
