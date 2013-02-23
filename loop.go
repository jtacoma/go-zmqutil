// Copyright 2012-2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gozmqutil

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

// A SocketEvent represents a set of events on a socket.
//
type SocketEvent struct {
	Socket  *Socket        // socket on which events occurred
	Events  zmq.PollEvents // bitmask of events that occurred
	Fault   error          // handlers may set this
	Message [][]byte       // when Events&POLLIN, the full message
}

// A SocketHandler acts on a *SocketEvent.
//
// When a SocketHandler returns an error to a poll loop, the loop will exit.  An
// instance of this interface is returned from *Loop.HandleFunc(), and can be
// passed to *Loop.Unhandle() to unsubscribe.
//
type SocketHandler interface {
	HandleSocketEvent(*SocketEvent)
}

type socketHandlerFunc struct {
	fun func(*SocketEvent)
}

func (h socketHandlerFunc) HandleSocketEvent(e *SocketEvent) {
	h.fun(e)
}

// A Loop is a ZeroMQ poll loop running in a goroutine.
//
// The loop will respond to events on sockets by calling handlers that have been
// associated with those events on those sockets through Handle() and
// HandleFunc().
//
// Note: since a Socket is not thread-safe, a Socket being polled by a Loop
// should not be operated on outside the scope of a handler.
//
type Loop struct {
	items   []loopItem  // sockets with their channels
	closing bool        // true if loop is either stopping or stopped
	logger  *log.Logger // changes when SetVerbose is called
	locker  *sync.Mutex // synchronize access to Loop state
}

type loopItem struct {
	socket  *Socket        // socket to poll
	events  zmq.PollEvents // events to poll for
	handler SocketHandler  // func to call when events occur
}

// NewLoop creates a new poll loop.
//
func NewLoop(context *Context) *Loop {
	return &Loop{
		locker: &sync.Mutex{},
	}
}

// Handle adds a socket event handler to p.
//
func (p *Loop) Handle(s *Socket, e zmq.PollEvents, h SocketHandler) {
	p.locker.Lock()
	defer p.locker.Unlock()
	var exists bool
	for _, existing := range p.items {
		if existing.socket == s && existing.handler == h {
			existing.events = existing.events | e
			exists = true
		}
	}
	if !exists {
		p.items = append(p.items, loopItem{s, e, h})
	}
}

// HandleFunc adds a socket event handler to p.
//
func (p *Loop) HandleFunc(s *Socket, e zmq.PollEvents, h func(*SocketEvent)) SocketHandler {
	handler := &socketHandlerFunc{h}
	p.Handle(s, e, handler)
	return handler
}

// Unhandle removes a handler from p for the given socket and socket events.
//
// If there are no remaining handlers for any event on this socket, the socket
// itself will cease to be polled in p.
//
func (p *Loop) Unhandle(s *Socket, e zmq.PollEvents, h SocketHandler) {
	p.locker.Lock()
	defer p.locker.Unlock()
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
}

// SetVerbose enables (or disables) logging to os.Stdout.
//
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

func (p *Loop) logf(s string, args ...interface{}) {
	if p.logger != nil {
		if s[len(s)-1] != '\n' {
			s += "\n"
		}
		p.logger.Printf("[gozmqutil] "+s, args...)
	}
}

// Run calls Step until an error is returned; then it closes p.
//
func (p *Loop) Run() error {
	for !p.closing {
		if err := p.Step(-1); err != nil {
			p.logf("loop: closing notification-receiving socket...")
			p.closing = true
			return err
		}
	}
	return nil
}

// Step polls for all handled events on whndled sockets for the
// specified timeout and then calls handlers for any detected events.
//
// A negative timeout means forever; otherwise, timeout wll be truncated
// to millisecond precision.
//
// Execution will halt and return first error encountered from polling
// or handling.
//
func (p *Loop) Step(timeout time.Duration) (err error) {

	// This PollItems construction may become inefficient for large
	// numbers of handlers.
	pollItems := make(zmq.PollItems, 0, len(p.items))
	for _, item := range p.items {
		var exists bool
		for _, existing := range pollItems {
			if existing.Socket == item.socket.base {
				existing.Events = existing.Events | item.events
				exists = true
			}
		}
		if !exists {
			pollItems = append(pollItems, zmq.PollItem{
				Socket: item.socket.base,
				Events: item.events,
			})
		}
	}

	// Truncate to the nearest
	ms := int64(timeout / time.Millisecond)
	if timeout < 0 {
		ms = -1
	}
	_, err = zmq.Poll(pollItems, ms)

	// Possible errors returned from Poll() are: ETERM, meaning a
	// context was closed; EFAULT, meaning a mistake was made in
	// setting up the PollItems list; and EINTR, meaning a signal
	// was delivered before any events were available.  Here, we
	// treat all errors the same:
	if err != nil {
		return err
	}

	p.logf("loop: events detected.")

	// Check all other sockets, sending any available messages to
	// their associated channels:
	for _, pollItem := range pollItems {
		event := SocketEvent{
			Events: pollItem.REvents,
		}
		for _, loopItem := range p.items {
			if loopItem.socket.base == pollItem.Socket && (loopItem.events&pollItem.REvents) != 0 {
				event.Socket = loopItem.socket
				if (pollItem.REvents&zmq.POLLIN) != 0 && event.Message == nil {
					event.Message, err = event.Socket.RecvMultipart(0)
					if err != nil {
						return err
					}
				}
				loopItem.handler.HandleSocketEvent(&event)
				if event.Fault != nil {
					return event.Fault
				}
			}
		}
	}

	return nil
}

// newPair returns a PUSH/PULL pair of inproc sockets.
func newPair(c *Context) (send *Socket, recv *Socket, err error) {
	send, err = c.NewSocket(zmq.PUSH)
	if err != nil {
		return
	}
	send.SetLinger(0)
	recv, err = c.NewSocket(zmq.PULL)
	if err != nil {
		return
	}
	send.SetLinger(0)
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

// newInprocAddress returns a unique incproc address.
func newInprocAddress() string {
	inprocNextMutex.Lock()
	defer inprocNextMutex.Unlock()
	inprocNext += 1
	return "inproc://github.com/jtacoma/gozmqutil/" + strconv.Itoa(inprocNext-1)
}

var inprocNext = 1
var inprocNextMutex sync.Mutex
