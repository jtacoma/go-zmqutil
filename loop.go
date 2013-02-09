// Copyright 2012-2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gzmq package lets messages be received from and sent to ZeroMQ
// sockets through channels.
package gzmq

// TODO: make loop.fault available to callers outside this package.

import (
	"errors"
	"strconv"
	"sync"

	zmq "github.com/alecthomas/gozmq"
)

// NewSending starts a goroutine that pumps messages from a channel to a
// socket.
//
// If messages will also be received from this socket, do not use this
// function.  Instead, use the corresponding method on a Loop.
//
// The goroutine will continue until the channel is closed or an error
// is encountered.  Any error other than ETERM will cause a panic.
func NewSending(s Socket) (chan<- [][]byte, error) {
	channel := make(chan [][]byte)
	go func() {
	F:
		for msg := range channel {
			err := s.SendMultipart(msg, 0)
			if err != nil {
				switch err {
				case zmq.ETERM:
					break F
				default:
					panic(err.Error())
				}
			}
		}
	}()
	return channel, nil
}

// A Loop is a ZeroMQ poll loop running in a goroutine.
type Loop interface {
	NewSending(Socket) (chan<- [][]byte, error)
	Start(Socket, int) (<-chan [][]byte, error)
	Stop(Socket) error
	Sync(func())
	Close() error
}

type loop struct {
	notifySend Socket         // to notify goroutine of pending commands
	commands   chan func()    // pending commands
	items      []loopItem     // sockets with their channels
	fault      error          // error, if any, that caused loop to stop
	closing    bool           // true if loop is either stopping or stopped
	running    sync.WaitGroup // becomes zero when loop is finished
}

type loopItem struct {
	socket  Socket        // socket that is/will be polled
	channel chan [][]byte // messages that have been received
}

// NewLoop starts a ZeroMQ poll loop in a goroutine.
//
// The cmdBuf argument specifies the length of the buffered channel that
// will receive commands to be executed inside the polling loop.
func NewLoop(context Context, cmdBuf int) (Loop, error) {
	notifySend, notifyRecv, err := newPair(context)
	if err != nil {
		return nil, err
	}
	p := loop{
		notifySend: notifySend,
		commands:   make(chan func(), cmdBuf),
		items: []loopItem{
			loopItem{notifyRecv, nil},
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
func (p *loop) Close() error {
	p.Sync(func() { p.closing = true })
	p.notifySend.Close()
	p.running.Wait()
	return nil
}

// NewSending starts a goroutine that pumps messages from a channel to a
// socket.
//
// The goroutine will continue until the channel is closed or an error
// is encountered.  Any error other than ETERM will cause a panic.
func (p *loop) NewSending(s Socket) (chan<- [][]byte, error) {
	channel := make(chan [][]byte)
	go func() {
		done := make(chan bool)
		for msg := range channel {
			p.Sync(func() {
				err := s.SendMultipart(msg, 0)
				if err != nil {
					switch err {
					case zmq.ETERM:
						done <- true
					default:
						panic(err.Error())
					}
				} else {
					done <- false
				}
			})
			if <-done {
				break
			}
		}
	}()
	return channel, nil
}

// Start adds a ZeroMQ socket to a poll loop, so that it will be
// polled for incoming messages.
//
// Notice that while this loop is running you must not use the socket
// in any other way except within the scope of a func passed to the Sync
// method.
//
// The chanBuf argument specifies the length of the buffered channel
// that will queue received messages for processing.
func (p *loop) Start(s Socket, chanBuf int) (result <-chan [][]byte, err error) {
	done := make(chan int)
	p.Sync(func() {
		var ok bool
		for _, existing := range p.items {
			if existing.socket == s {
				ok = true
			}
		}
		if !ok {
			ch := make(chan [][]byte, chanBuf)
			p.items = append(p.items, loopItem{s, ch})
			result = ch
		}
		done <- 1
	})
	<-done
	return
}

// Stop removes a ZeroMQ socket from a poll loop, so that it will no
// longer be polled for incoming messages.
func (p *loop) Stop(s Socket) (err error) {
	done := make(chan int)
	p.Sync(func() {
		index := -1
		for i, existing := range p.items {
			if existing.socket == s {
				index = i
			}
		}
		if index < 0 {
			err = errors.New("socket is already not being polled.")
		} else if index == len(p.items)-1 {
			p.items = p.items[0:index]
		} else {
			p.items = append(p.items[0:index], p.items[index+1:len(p.items)]...)
		}
		done <- 1
	})
	<-done
	return err
}

// Sync executes a function outside of the poll.
//
// Useful for performing operations on the sockets being polled.
func (p *loop) Sync(f func()) {
	p.notifySend.Send([]byte{0}, 0)
	p.commands <- f
}

func (p *loop) loop(notifyRecv Socket) {

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

		// Check for incoming commands.  For each message available in
		// notifyRecv, dequeue one command and execute it.
		//
		// Commands may modify p.items, which earlier code in this
		// method assumes aligns with local slice pollItems.  Therfore,
		// commands must be processed afterward.
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
	}
}

func newPair(c Context) (send Socket, recv Socket, err error) {
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
