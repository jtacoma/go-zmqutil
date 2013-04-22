// Copyright 2012-2013 Joshua Tacoma
//
// This file is part of ZeroMQ Utilities.
//
// ZeroMQ Utilities is free software: you can redistribute it and/or
// modify it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// ZeroMQ Utilities is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public
// License along with ZeroMQ Utilities.  If not, see
// <http://www.gnu.org/licenses/>.

package zmqutil

import (
	"log"
	"sync"
	"syscall"
	"time"

	zmq "github.com/alecthomas/gozmq"
)

// A Event represents a set of events on a socket.
//
type Event struct {
	Socket *Socket        // socket on which events occurred
	Events zmq.PollEvents // bitmask of events that occurred
	Fault  error          // handlers may set this to halt the poller
}

// A Poller is a ZeroMQ poller running in a goroutine.
//
// The poller will respond to events on sockets by calling handlers that have
// been associated with those events on those sockets through Handle() and
// HandleFunc().
//
// Note: since a Socket is not thread-safe, a Socket being polled by a Poller
// should not be operated on outside the scope of a handler.
//
type Poller struct {
	items  map[*zmq.Socket]*pollItem // sockets with their channels
	locker *sync.Mutex               // synchronize access to Poller state
	logger *log.Logger
}

type pollItem struct {
	socket    *Socket        // socket to poll
	events    zmq.PollEvents // events to poll for
	handleErr func(error)    // func to call when an error occurs
	handleIn  func([][]byte) // func to call when messages arrive
	handleOut func()         // func to call when messages can be sent
}

// NewPoller creates a new poller.
//
func NewPoller(context *Context) *Poller {
	return &Poller{
		items:  make(map[*zmq.Socket]*pollItem),
		locker: &sync.Mutex{},
	}
}

// HandleErr sets a function to be called when an error occurs on s.
//
func (p *Poller) HandleErr(s *Socket, h func()) {
	p.locker.Lock()
	defer p.locker.Unlock()
	item, ok := p.items[s.s]
	if !ok {
		item = &pollItem{socket: s}
		p.items[s.s] = item
	}
	if h == nil {
		item.events = item.events ^ zmq.POLLOUT
	} else {
		item.events = item.events | zmq.POLLOUT
	}
	if item.events == 0 {
		p.Unhandle(s)
	} else {
		item.handleOut = h
	}
}

// HandleIn sets the function that will be called for each message that
// arrives on s.
//
func (p *Poller) HandleIn(s *Socket, h func([][]byte)) {
	p.locker.Lock()
	defer p.locker.Unlock()
	item, ok := p.items[s.s]
	if !ok {
		item = &pollItem{socket: s}
		p.items[s.s] = item
	}
	if h == nil {
		item.events &= ^zmq.POLLIN
	} else {
		item.events |= zmq.POLLIN
	}
	if item.events == 0 {
		p.Unhandle(s)
	} else {
		item.handleIn = h
	}
}

// HandleOut sets the function that will be called when a message can be sent
// on s with no delay.
//
func (p *Poller) HandleOut(s *Socket, h func()) {
	p.locker.Lock()
	defer p.locker.Unlock()
	item, ok := p.items[s.s]
	if !ok {
		item = &pollItem{socket: s}
		p.items[s.s] = item
	}
	if h == nil {
		item.events = item.events ^ zmq.POLLOUT
	} else {
		item.events = item.events | zmq.POLLOUT
	}
	if item.events == 0 {
		p.Unhandle(s)
	} else {
		item.handleOut = h
	}
}

// Unhandle removes any handlers for s and stops polling s.
//
func (p *Poller) Unhandle(s *Socket) {
	p.locker.Lock()
	defer p.locker.Unlock()
	if _, ok := p.items[s.s]; ok {
		delete(p.items, s.s)
	}
}

// SetLogger sets the logger that detailed messages will be sent to.
//
func (p *Poller) SetLogger(logger *log.Logger) {
	p.logger = logger
	p.logf("logger received.")
}

func (p *Poller) logf(s string, args ...interface{}) {
	if p.logger != nil {
		if s[len(s)-1] != '\n' {
			s += "\n"
		}
		p.logger.Printf("[zmqutil] "+s, args...)
	}
}

// Run repeatedly calls Poll with an infinite timeout until an error is
// returned, then returns that error.
//
func (p *Poller) Run() error {
	for {
		if err := p.Poll(-1); err != nil {
			return err
		}
	}
	return nil
}

// Poll polls, with the specified timeout, all sockets for all events that have
// been registered with event handlers.
//
// A negative timeout means forever; otherwise, timeout wll be truncated to
// millisecond precision.
//
// Execution will halt and return first error encountered from polling
// or handling.
//
func (p *Poller) Poll(timeout time.Duration) (err error) {
	p.locker.Lock()
	defer p.locker.Unlock()

	// This PollItems construction may become inefficient for large
	// numbers of handlers.
	baseItems := make(zmq.PollItems, 0, len(p.items))
	for s, item := range p.items {
		baseItems = append(baseItems, zmq.PollItem{
			Socket: s,
			Events: item.events,
		})
	}

	p.logf("poller: polling %d sockets for %s", len(baseItems), timeout)
	n, err := zmq.Poll(baseItems, timeout)

	// Possible errors returned from Poll() are: ETERM, meaning a
	// context was closed; EFAULT, meaning a mistake was made in
	// setting up the PollItems list; and EINTR, meaning a signal
	// was delivered before any events were available.  Here, we
	// treat all errors the same:
	if err != nil {
		p.logf("poller: error while polling: %s", err)
		return err
	}

	if n > 0 {
		p.logf("poller: events detected.")

		// Check all other sockets, sending any available messages to
		// their associated channels:
		for _, base := range baseItems {
			item := p.items[base.Socket]
			if (base.Events&zmq.POLLIN) != 0 && item.handleIn != nil {
				for {
					m, err := base.Socket.RecvMultipart(zmq.DONTWAIT)
					if err == syscall.EAGAIN {
						break
					} else if err != nil {
						if item.handleErr != nil {
							item.handleErr(err)
						}
						break
					}
					item.handleIn(m)
				}
			}
			if (base.Events&zmq.POLLOUT) != 0 && item.handleOut != nil {
				item.handleOut()
			}
		}
	}
	return nil
}
