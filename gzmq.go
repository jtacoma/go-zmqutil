package gzmq

import (
	"strconv"
	"sync"

	zmq "github.com/alecthomas/gozmq"
)

type Polling interface {
	Include(zmq.Socket) (<-chan [][]byte, error)
	Close() error
}

type polling struct {
	notifySend zmq.Socket  // to notify goroutine of pending commands
	commands   chan func() // pending commands
	items      []pollingItem
	closing    bool
	closed     sync.WaitGroup
}

type pollingItem struct {
	socket  zmq.Socket
	channel chan [][]byte
}

func NewPolling(context zmq.Context) (Polling, error) {
	notifySend, notifyRecv, err := newPair(context)
	if err != nil {
		return nil, err
	}
	p := polling{
		notifySend: notifySend,
		commands:   make(chan func(), 4),
		items: []pollingItem{
			pollingItem{notifyRecv, nil},
		},
	}
	p.closed.Add(1)
	go func() {
		for !p.closing {
			pollItems := make(zmq.PollItems, len(p.items))
			for i, item := range p.items {
				pollItems[i].Socket = item.socket
				pollItems[i].Events = zmq.POLLIN
			}
			_, err := zmq.Poll(pollItems, -1)
			if err != nil {
				println("E:", err.Error())
				break
			}
			if (pollItems[0].REvents & zmq.POLLIN) != 0 {
				_, err := notifyRecv.RecvMultipart(0)
				if err != nil {
					println("E:", err.Error())
					break
				}
				cmd := <-p.commands
				cmd()
			}
			for i := 1; i < len(pollItems); i++ {
				item := pollItems[i]
				if (item.REvents & zmq.POLLIN) != 0 {
					msg, err := item.Socket.RecvMultipart(0)
					if err != nil {
						println("E:", err.Error())
						continue //?
					}
					pitem := p.items[i]
					println("I: receiving message into channel.")
					pitem.channel <- msg
					println("I: received message into channel.")
				}
			}
		}
		notifyRecv.Close()
		p.closed.Done()
	}()
	return &p, nil
}

func (p *polling) Close() error {
	p.notifySend.Send([]byte{0}, 0)
	p.notifySend.Close()
	p.commands <- func() { p.closing = true }
	p.closed.Wait()
	return nil
}

func (p *polling) Include(s zmq.Socket) (result <-chan [][]byte, err error) {
	done := make(chan int)
	p.notifySend.Send([]byte{0}, 0)
	p.commands <- func() {
		var ok bool
		for _, existing := range p.items {
			if existing.socket == s {
				ok = true
			}
		}
		if !ok {
			ch := make(chan [][]byte, 4)
			println("I: including socket", s)
			p.items = append(p.items, pollingItem{s, ch})
			result = ch
		}
		done <- 1
	}
	<-done
	return
}

func newPair(c zmq.Context) (send zmq.Socket, recv zmq.Socket, err error) {
	send, err = c.NewSocket(zmq.PUSH)
	if err != nil {
		return
	}
	recv, err = c.NewSocket(zmq.PULL)
	if err != nil {
		return
	}
	addr := newInprocAddress()
	err = send.Bind(addr)
	if err != nil {
		return
	}
	err = recv.Connect(addr)
	if err != nil {
		return
	}
	return
}

var inprocNext = 1

func newInprocAddress() string {
	inprocNext += 1
	return "inproc://github.com/jtacoma/gzmq/" + strconv.Itoa(inprocNext-1)
}
