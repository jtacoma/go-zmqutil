// Copyright 2012-2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzmq

// TODO: make loop.fault available to callers outside this package.

import (
	"log"
	"os"
	"strconv"
	"sync"

	zmq "github.com/alecthomas/gozmq"
)

// A SocketEvent represents an event on a socket.
type SocketEvent struct {
	Socket Socket
	Events zmq.PollEvents
}

// A SocketHandler is a comparable wrapper for what is really just a func.
type SocketHandler struct {
	fun func(*SocketEvent) error
}

// A Loop is a ZeroMQ poll loop running in a goroutine.
type Loop struct {
	notifySend Socket         // to notify goroutine of pending commands
	commands   chan func()    // pending commands
	items      []loopItem     // sockets with their channels
	fault      error          // error, if any, that caused loop to stop
	closing    bool           // true if loop is either stopping or stopped
	running    sync.WaitGroup // becomes zero when loop is finished
	logger     *log.Logger
}

type loopItem struct {
	socket  Socket         // socket to poll
	events  zmq.PollEvents // events to poll for
	handler *SocketHandler // func to call when events occur
}

// NewLoop starts a poll loop in a goroutine.
//
// The cmdBuf argument specifies the length of the buffered channel that
// will receive commands to be executed inside the polling loop.
func NewLoop(context Context) (*Loop, error) {
	notifySend, notifyRecv, err := newPair(context)
	if err != nil {
		return nil, err
	}
	p := Loop{
		notifySend: notifySend,
		commands:   make(chan func(), 64),
		items: []loopItem{
			loopItem{notifyRecv, zmq.POLLIN, nil},
		},
	}
	p.running.Add(1)
	go p.loop(notifyRecv)
	return &p, nil
}

// Close causes the poll loop to exit, and blocks until this is done.
//
// This will not close polled sockets.  The loop cannot be re-opened.
func (p *Loop) Close() error {
	p.logf("loop: enqueuing func to mark as closing...")
	p.Sync(func() { p.closing = true })
	p.logf("loop: closing the notification-sending socket...")
	p.notifySend.Close()
	p.logf("loop: waiting for loop to stop runnning...")
	p.running.Wait()
	p.logf("loop: stopped.")
	return p.fault
}

// PollSocket adds a socket to a poll loop.
//
// Notice that while this loop is running you must not use the socket
// in any other way except within the scope of a func passed to the Sync
// method.
func (p *Loop) PollSocket(s Socket, e zmq.PollEvents, h func(*SocketEvent) error) (handler *SocketHandler, err error) {
	done := make(chan int)
	handler = &SocketHandler{h}
	p.Sync(func() {
		var exists bool
		for _, existing := range p.items {
			if existing.socket == s && existing.handler == handler {
				existing.events = existing.events | e
				exists = true
			}
		}
		if !exists {
			p.items = append(p.items, loopItem{s, e, handler})
		}
		done <- 1
	})
	<-done
	return
}

// Stop removes a ZeroMQ socket from a poll loop, so that it will no
// longer be polled for incoming messages.
func (p *Loop) PollSocketEnd(s Socket, e zmq.PollEvents, h *SocketHandler) (err error) {
	done := make(chan int)
	p.Sync(func() {
		index := -1
		for i, existing := range p.items {
			if existing.socket == s && existing.handler == h {
				existing.events = existing.events & ^e
				if existing.events == 0 {
					index = i
				}
			}
		}
		if index == len(p.items)-1 {
			p.items = p.items[0:index]
		} else if index >= 0 {
			p.items = append(p.items[0:index], p.items[index+1:len(p.items)]...)
		}
		done <- 1
	})
	<-done
	return err
}

func (p *Loop) SetVerbose(verbose bool) error {
	if verbose == (p.logger != nil) {
		return nil
	}
	p.logf("loop.verbose = %t", verbose)
	if verbose && p.logger == nil {
		p.logger = log.New(os.Stdout, "", log.Lmicroseconds)
	} else if !verbose {
		p.logger = nil
	}
	p.logf("loop.verbose = %t", p.logger != nil)
	return nil
}

// Sync executes a function outside of the poll.
//
// Useful for performing operations on the sockets being polled.
func (p *Loop) Sync(f func()) {
	p.notifySend.Send([]byte{0}, 0)
	p.commands <- f
}

func (p *Loop) logf(s string, args ...interface{}) {
	if p.logger != nil {
		if s[len(s)-1] != '\n' {
			s += "\n"
		}
		p.logger.Printf("[gzmq] "+s, args...)
	}
}

func (p *Loop) loop(notifyRecv Socket) {

	defer func() {
		p.logf("loop: closing notification-receiving socket...")
		notifyRecv.Close()
		p.logf("loop: declaring that loop is no longer running...")
		p.running.Done()
	}()

	for !p.closing {

		// TODO: refactor so that pollItems is re-constructed only when
		// necessary.
		pollItems := make(zmq.PollItems, 0, len(p.items))
		for _, item := range p.items {
			var exists bool
			for _, existing := range pollItems {
				if existing.Socket == item.socket {
					existing.Events = existing.Events | item.events
					exists = true
				}
			}
			if !exists {
				pollItems = append(pollItems, zmq.PollItem{
					Socket: item.socket,
					Events: item.events,
				})
			}
		}

		// Poll with an infinite timeout: this loop never spins idle.
		_, err := zmq.Poll(pollItems, -1)

		// Possible errors returned from Poll() are: ETERM, meaning a
		// context was closed; EFAULT, meaning a mistake was made in
		// setting up the PollItems list; and EINTR, meaning a signal
		// was delivered before any events were available.  Here, we
		// treat all errors the same:
		if err != nil {
			p.logf("loop: error while polling: %s", err.Error())
			p.fault = err
			p.closing = true
			break
		}

		p.logf("loop: events detected.")

		// Check all other sockets, sending any available messages to
		// their associated channels:
		for i := 1; i < len(pollItems); i++ {
			pollItem := pollItems[i]
			event := SocketEvent{
				Socket: pollItem.Socket,
				Events: pollItem.REvents,
			}
			for _, loopItem := range p.items {
				if loopItem.socket == pollItem.Socket && (loopItem.events&pollItem.REvents) != 0 {
					if err = loopItem.handler.fun(&event); err != nil {
						p.fault = err
						p.closing = true
						continue //?
					}
				}
			}
		}

		// Check for incoming commands.  For each message available in
		// notifyRecv, dequeue one command and execute it.
		//
		// Commands may modify p.items, which earlier code in this
		// method assumes aligns with local slice pollItems.  Therefore,
		// commands must be processed afterward.  (TODO: this clause is
		// no longer true, roll command handling into the rest!)
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
	send.SetSockOptInt(zmq.LINGER, 0)
	recv, err = c.NewSocket(zmq.PULL)
	if err != nil {
		return
	}
	send.SetSockOptInt(zmq.LINGER, 0)
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
