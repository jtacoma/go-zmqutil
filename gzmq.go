// The gzmq package provides abstractions that make ZeroMQ more
// accessible to idiomatic Go.
//
// ZeroMQ's sockets and Go's channels are similar in purpose, but have
// some significant differences.  ZeroMQ's sockets can communicate
// across process and machine boundaries, while Go's channels cannot.
// Go's channels are thread-safe, while ZeroMQ's sockets are not.
//
// ZeroMQ's Poll() method and Go's "select" statement are attempts to
// solve the same problem: efficient reads from multiple sources of
// information.  However, ZeroMQ's Poll() method cannot be used to poll
// a combination of sockets and channels, while Go's "select" statement
// requires explicit code blocks for each information source.
// 
// The Polling defined in this package is an attempt to makee ZeroMQ sockets
// available, through a Poll() loop, as channels for use in Go "select"
// statements.
package gzmq

// TODO: make polling.fault available to callers outside this package.

import (
	"strconv"
	"sync"

	zmq "github.com/alecthomas/gozmq"
)

// A Polling is a ZeroMQ poll loop running in a goroutine.
type Polling interface {
	Include(zmq.Socket) (<-chan [][]byte, error)
	Sync(func())
	Close() error
}

type polling struct {
	notifySend zmq.Socket     // to notify goroutine of pending commands
	commands   chan func()    // pending commands
	items      []pollingItem  // sockets with their channels
	fault      error          // error, if any, that caused polling to stop
	closing    bool           // true if polling is either stopping or stopped
	running    sync.WaitGroup // becomes zero when polling is finished
}

type pollingItem struct {
	socket  zmq.Socket    // socket that is/will be polled
	channel chan [][]byte // messages that have been received
}

// NewPolling starts a ZeroMQ poll loop in a goroutine.
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
	p.running.Add(1)
	go p.loop(notifyRecv)
	return &p, nil
}

// Close causes the poll loop to exit, and blocks until it is done.
//
// Message channels will be closed, but polled sockets will be left
// open.
func (p *polling) Close() error {
	p.Sync(func() { p.closing = true })
	p.notifySend.Close()
	p.running.Wait()
	return nil
}

// Include adds a ZeroMQ socket to a poll loop, so that it will be
// polled for incoming messages.
//
// Notice that while this polling is running you must not use the socket
// in any other way except within the scope of a func passed to the Lock
// method.
//
// The returned channel must be used to receive messages from this
// socket until the polling is closed.
func (p *polling) Include(s zmq.Socket) (result <-chan [][]byte, err error) {
	done := make(chan int)
	p.Sync(func() {
		var ok bool
		for _, existing := range p.items {
			if existing.socket == s {
				ok = true
			}
		}
		if !ok {
			ch := make(chan [][]byte, 4)
			p.items = append(p.items, pollingItem{s, ch})
			result = ch
		}
		done <- 1
	})
	<-done
	return
}

// Sync executes a function outside of the poll.
//
// Useful for performing operations on the sockets being polled.
func (p *polling) Sync(f func()) {
	p.notifySend.Send([]byte{0}, 0)
	p.commands <- f
}

func (p *polling) loop(notifyRecv zmq.Socket) {

	defer func() {
		// Close all "received message" channels :
		for _, pi := range p.items {
			if pi.channel != nil {
				close(pi.channel)
			}
		}

		// For sockets that may have queued out-going messages, this call
		// may block until they are sent:
		notifyRecv.Close()

		// Indicate to any waiting code that this poll loop is all done:
		p.running.Done()
	}()

	for !p.closing {

		// TODO: refactor so that pollItems is re-constructed only when
		// necessary.
		pollItems := make(zmq.PollItems, len(p.items))
		for i, item := range p.items {
			pollItems[i].Socket = item.socket
			pollItems[i].Events = zmq.POLLIN
		}

		// Poll with an infinite timeout: this loop never spins idle.
		_, err := zmq.Poll(pollItems, -1)

		// Possible errors returned from Poll() are: ETERM, meaning a
		// context was closed; EFAULT, meaning a mistake was made in
		// setting up the PollItems list; and EINTR, meaning a signal
		// was delivered before any events were available.  Here, we
		// treat all errors the same:
		if err != nil {
			p.fault = err
			p.closing = true
			break
		}

		// Check for incoming commands first.  For each message
		// available in notifyRecv, dequeue one command and execute it:
		if (pollItems[0].REvents & zmq.POLLIN) != 0 {
			_, err := notifyRecv.RecvMultipart(0)
			if err != nil {
				p.fault = err
				p.closing = true
				break
			}
			cmd := <-p.commands
			cmd()
		}

		// Check all other sockets, sending any available messages to
		// their associated channels:
		for i := 1; i < len(pollItems); i++ {
			item := pollItems[i]
			if (item.REvents & zmq.POLLIN) != 0 {
				msg, err := item.Socket.RecvMultipart(0)
				if err != nil {
					p.fault = err
					p.closing = true
					continue //?
				}
				pitem := p.items[i]
				pitem.channel <- msg
			}
		}
	}
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
