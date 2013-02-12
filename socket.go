// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gozmqutil

import (
	"time"

	zmq "github.com/alecthomas/gozmq"
)

type Socket struct {
	base zmq.Socket
	ctx  *Context
}

func (s *Socket) Close() error {
	return s.base.Close()
}

func (s *Socket) Bind(addr string) error    { return s.base.Bind(addr) }
func (s *Socket) Connect(addr string) error { return s.base.Connect(addr) }

func (s *Socket) Recv(flags zmq.SendRecvOption) ([]byte, error) {
	return s.base.Recv(flags)
}
func (s *Socket) RecvMultipart(flags zmq.SendRecvOption) ([][]byte, error) {
	return s.base.RecvMultipart(flags)
}
func (s *Socket) Send(frame []byte, flags zmq.SendRecvOption) error {
	return s.base.Send(frame, flags)
}
func (s *Socket) SendMultipart(msg [][]byte, flags zmq.SendRecvOption) error {
	return s.base.SendMultipart(msg, flags)
}

func (s *Socket) Linger() (time.Duration, error) {
	if s == nil {
		return -1, SocketIsNil
	}
	ms, err := s.base.GetSockOptInt(zmq.LINGER)
	if err != nil {
		return -1, err
	}
	if ms < 0 {
		return -1, nil
	}
	return time.Duration(ms) * time.Millisecond, nil
}

// Sets the high water mark for outbound messages on the specified socket.
//
// The high water mark is a hard limit on the maximum number of outstanding
// messages ØMQ shall queue in memory for any single peer that the specified
// socket is communicating with.
//
// If this limit has been reached the socket shall enter an exceptional state
// and depending on the socket type, ØMQ shall take appropriate action such as
// blocking or dropping sent messages.  Refer to the individual socket
// descriptions in zmq_socket(3) for details on the exact action taken for each
// socket type.
//
// ØMQ does not guarantee that the socket will accept this many messages, and
// the actual limit may be as much as 60-70% lower depending on the flow of
// messages on the socket.
//
func (s *Socket) SetSendHWM(int) error { return NotImplemented }

// Sets the high water mark for inbound messages on the specified socket.
//
// The high water mark is a hard limit on the maximum number of outstanding
// messages ØMQ shall queue in memory for any single peer that the specified
// socket is communicating with.
//
// If this limit has been reached the socket shall enter an exceptional state
// and depending on the socket type, ØMQ shall take appropriate action such as
// blocking or dropping sent messages.  Refer to the individual socket
// descriptions in zmq_socket(3) for details on the exact action taken for each
// socket type.
//
func (s *Socket) SetRecvHWM(int) error { return NotImplemented }

// Sets the I/O thread affinity for newly created connections on the specified
// socket.
//
// Affinity determines which threads from the ØMQ I/O thread pool associated
// with the socket's context shall handle newly created connections.  A value of
// zero specifies no affinity, meaning that work shall be distributed fairly
// among all ØMQ I/O threads in the thread pool. For non-zero values, the lowest
// bit corresponds to thread 1, second lowest bit to thread 2 and so on.  For
// example, a value of 3 specifies that subsequent connections on socket shall
// be handled exclusively by I/O threads 1 and 2.
//
// See also zmq_init(3) for details on allocating the number of I/O threads for
// a specific context.
//
func (s *Socket) SetAffinity(threadmask uint64) error { return NotImplemented }

// Establishs a new message filter on a SUB socket.
//
// Newly created SUB sockets shall filter out all incoming messages, therefore
// you should call this option to establish an initial message filter.
//
// An empty option_value of length zero shall subscribe to all incoming
// messages.  A non-empty option_value shall subscribe to all messages beginning
// with the specified prefix.  Multiple filters may be attached to a single SUB
// socket, in which case a message shall be accepted if it matches at least one
// filter.
//
func (s *Socket) Subscribe([]byte) error { return NotImplemented }

// Removes an existing message filter on a SUB socket.
//
// The filter specified must match an existing filter previously established
// with the ZMQ_SUBSCRIBE option.  If the socket has several instances of the
// same filter attached only one instance will be removed, leaving the rest in
// place and functional.
//
func (s *Socket) Unsubscribe([]byte) error { return NotImplemented }

// Set the identity of the specified socket.
//
// Socket identity is used only by request/reply pattern.  Namely, it can be
// used in tandem with ROUTER socket to route messages to the peer with specific
// identity.
//
// Identity should be at least one byte and at most 255 bytes long.  Identities
// starting with binary zero are reserved for use by ØMQ infrastructure.
//
// If two peers use the same identity when connecting to a third peer, the
// results shall be undefined.
//
func (s *Socket) SetIdentity(string) error { return NotImplemented }

// Sets the maximum send or receive data rate (in kilobits per second) for
// multicast transports such as zmq_pgm(7) using the specified socket.
//
func (s *Socket) SetRate(kbps int) error { return NotImplemented }

// Sets the recovery interval for multicast transports using the specified
// socket.
//
// The recovery interval determines the maximum time in milliseconds that a
// receiver can be absent from a multicast group before unrecoverable data loss
// will occur.
//
// Exercise care when setting large recovery intervals as the data needed for
// recovery will be held in memory.  For example, a 1 minute recovery interval
// at a data rate of 1Gbps requires a 7GB in-memory buffer.
func (s *Socket) SetRecoveryInterval(time.Duration) error { return NotImplemented }

// Sets the underlying kernel transmit buffer size for the socket to the
// specified size in bytes.
//
// A value of zero means leave the OS default unchanged.  For details please
// refer to your operating system documentation for the SO_SNDBUF socket option.
//
func (s *Socket) SetSendBuf(bytes int) error { return NotImplemented }

// Sets the underlying kernel receive buffer size for the socket to the
// specified size in bytes.
//
// A value of zero means leave the OS default unchanged.  For details refer to
// your operating system documentation for the SO_RCVBUF socket option.
//
func (s *Socket) SetRecvBuf(bytes int) error { return NotImplemented }

// Sets the linger period for the specified socket.
//
// The linger period determines how long pending messages which have yet to be
// sent to a peer shall linger in memory after a socket is closed, and further
// affects the termination of the socket's context.  The following outlines the
// different behaviours:
//
//  * The default value of -1 specifies an infinite linger period. Pending
//    messages shall not be discarded after a call to zmq_close(); attempting to
//    terminate the socket's context with zmq_term() shall block until all
//    pending messages have been sent to a peer.
//  * The value of 0 specifies no linger period. Pending messages shall be
//    discarded immediately when the socket is closed.
//  * Positive values, truncated to millisecond precision, specify an upper
//    bound for the linger period.  Pending messages shall not be discarded
//    after a socket is closed; attempting to terminate the socket's context
//    shall block until either all pending messages have been sent to a peer, or
//    the linger period expires, after which any pending messages shall be
//    discarded.
//
func (s *Socket) SetLinger(linger time.Duration) error {
	if s == nil {
		return SocketIsNil
	}
	var ms int
	if linger < 0 {
		ms = -1
	} else {
		ms = int(linger / time.Millisecond)
	}
	return s.base.SetSockOptInt(zmq.LINGER, ms)
}

// Sets the initial reconnection interval for the specified socket.
//
// The reconnection interval is the period ØMQ shall wait between attempts to
// reconnect disconnected peers when using connection-oriented transports.  The
// value -1 means no reconnection.
//
func (s *Socket) SetReconnectInterval(time.Duration) error { return NotImplemented }

// Sets the maximum reconnection interval for the specified socket.
//
// This is the maximum period ØMQ shall wait between attempts to reconnect.  On
// each reconnect attempt, the previous interval shall be doubled untill this
// maximum is reached.  This allows for exponential backoff strategy.  Default
// value means no exponential backoff is performed and reconnect interval
// calculations are only based on ReconnectInterval.
//
func (s *Socket) SetReconnectMaxInterval(time.Duration) error { return NotImplemented }

// Sets the maximum length of the queue of outstanding peer connections for the
// specified socket; this only applies to connection-oriented transports.
//
// For details refer to your operating system documentation for the listen function.
//
func (s *Socket) SetBacklog(connections int) error { return NotImplemented }

// Limits the size of the inbound message.
//
// If a peer sends a message larger than ZMQ_MAXMSGSIZE it is disconnected.
// Value of -1 means no limit.
//
func (s *Socket) SetMaxMsgSize(bytes int) error { return NotImplemented }

// Sets the time-to-live field in every multicast packet sent from this socket.
//
// The default is 1 which means that the multicast packets don't leave the local
// network.
//
func (s *Socket) SetMulticastHops(hops int) error { return NotImplemented }

// Sets the timeout for receive operation on the socket.
//
// If the value is 0, zmq_recv(3) will return immediately, with a EAGAIN error
// if there is no message to receive.  If the value is -1, it will block until a
// message is available.  For all other values, it will wait for a message for
// that amount of time before returning with an EAGAIN error.
//
func (s *Socket) SetRecvTimeout(time.Duration) error { return NotImplemented }

// Sets the timeout for send operation on the socket.
//
// If the value is 0, zmq_send(3) will return immediately, with a EAGAIN error
// if the message cannot be sent.  If the value is -1, it will block until the
// message is sent.  For all other values, it will try to send the message for
// that amount of time before returning with an EAGAIN error.
//
func (s *Socket) SetSendTimeout(time.Duration) error { return NotImplemented }

// Sets the underlying native socket type.
//
// A value of false will use IPv4 sockets, while the value of true will use IPv6
// sockets.  An IPv6 socket lets applications connect to and accept connections
// from both IPv4 and IPv6 hosts.
//
func (s *Socket) SetIPv6(bool) error { return NotImplemented }

// If set to true, will delay the attachment of a pipe on connect until the
// underlying connection has completed.
//
// This will cause the socket to block if there are no other connections, but
// will prevent queues from filling on pipes awaiting connection.
//
func (s *Socket) SetDelayAttachOnConnect(bool) error { return NotImplemented }

// Sets the XPUB socket behavior on new subscriptions and unsubscriptions.
//
// A value of false is the default and passes only new subscription messages to
// upstream.  A value of true passes all subscription messages upstream.
//
func (s *Socket) SetXPUBVerbose(bool) error { return NotImplemented }

// Sets the ROUTER socket behavior when an unroutable message is encountered.
//
// A value of 0 is the default and discards the message silently when it cannot
// be routed.  A value of 1 returns an EHOSTUNREACH error code if the message
// cannot be routed.
//
func (s *Socket) SetRouterMandatory(bool) error { return NotImplemented }

// Override SO_KEEPALIVE socket option (where supported by OS).
//
func (s *Socket) SetTCPKeepAlive(bool) error { return NotImplemented }

// Skip any overrides on SO_KEEPALIVE and leave it to OS default.
//
func (s *Socket) ResetTCPKeepAlive() error { return NotImplemented }

// Override TCP_KEEPCNT (or TCP_KEEPALIVE on some OS) socket option (where
// supported by OS).
//
// The default value of -1 means to skip any overrides and leave it to OS
// default.
//
func (s *Socket) SetTCPKeepAliveIdle(probes int) error { return NotImplemented }

// Override TCP_KEEPCNT socket option (where supported by OS).
//
// The default value of -1 means to skip any overrides and leave it to OS
// default.
func (s *Socket) SetTCPKeepAliveCount(probes int) error { return NotImplemented }

// Override TCP_KEEPINTVL socket option (where supported by OS).
//
// The default value of -1 means to skip any overrides and leave it to OS
// default.
func (s *Socket) SetTCPKeepAliveInterval(time.Duration) error { return NotImplemented }

// Assign arbitrary number of filters that will be applied for each new TCP
// transport connection on a listening socket.
//
// If no filters applied, then TCP transport allows connections from any ip.  If
// at least one filter is applied then new connection source ip should be
// matched.  To clear all filters call zmq_setsockopt(socket,
// ZMQ_TCP_ACCEPT_FILTER, NULL, 0). Filter is a null-terminated string with ipv6
// or ipv4 CIDR.
//
func (s *Socket) AddTCPAcceptFilter(string) error { return NotImplemented }
