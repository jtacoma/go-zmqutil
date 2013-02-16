// +build zmq_3_x
//
// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was generated automatically using https://gist.github.com/4969086.git

package gozmqutil

import (
	"time"

	zmq "github.com/alecthomas/gozmq"
)

// ZMQ_SNDHWM: Set high water mark for outbound messages.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_SNDHWM option shall set the high water mark for outbound messages on
// the specified socket. The high water mark is a hard limit on the maximum
// number of outstanding messages ØMQ shall queue in memory for any single peer
// that the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional state
// and depending on the socket type, ØMQ shall take appropriate action such as
// blocking or dropping sent messages. Refer to the individual socket
// descriptions in zmq_socket(3) for details on the exact action taken for each
// socket type.
// 
// ØMQ does not guarantee that the socket will accept as many as ZMQ_SNDHWM
// messages, and the actual limit may be as much as 60-70% lower depending on
// the flow of messages on the socket.
// 
//  Option value type         int
//  Option value unit         messages
//  Default value             1000
//  Applicable socket types   all
//
func (s *Socket) SetSndHWM(value int) error {
	return s.base.SetSockOptInt(zmq.SNDHWM, value)
}

// ZMQ_RCVHWM: Set high water mark for inbound messages.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_RCVHWM option shall set the high water mark for inbound messages on
// the specified socket. The high water mark is a hard limit on the maximum
// number of outstanding messages ØMQ shall queue in memory for any single peer
// that the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional state
// and depending on the socket type, ØMQ shall take appropriate action such as
// blocking or dropping sent messages. Refer to the individual socket
// descriptions in zmq_socket(3) for details on the exact action taken for each
// socket type.
// 
//  Option value type         int
//  Option value unit         messages
//  Default value             1000
//  Applicable socket types   all
//
func (s *Socket) SetRcvHWM(value int) error {
	return s.base.SetSockOptInt(zmq.RCVHWM, value)
}

// ZMQ_AFFINITY: Set I/O thread affinity.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_AFFINITY option shall set the I/O thread affinity for newly created
// connections on the specified socket.
// 
// Affinity determines which threads from the ØMQ I/O thread pool associated
// with the socket's context shall handle newly created connections. A value of
// zero specifies no affinity, meaning that work shall be distributed fairly
// among all ØMQ I/O threads in the thread pool. For non-zero values, the lowest
// bit corresponds to thread 1, second lowest bit to thread 2 and so on. For
// example, a value of 3 specifies that subsequent connections on socket shall
// be handled exclusively by I/O threads 1 and 2.
// 
// See also zmq_init(3) for details on allocating the number of I/O threads for
// a specific context.
// 
//  Option value type         uint64_t
//  Option value unit         N/A (bitmap)
//  Default value             0
//  Applicable socket types   N/A
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.AFFINITY, value)
}

// ZMQ_SUBSCRIBE: Establish message filter.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_SUBSCRIBE option shall establish a new message filter on a ZMQ_SUB
// socket. Newly created ZMQ_SUB sockets shall filter out all incoming messages,
// therefore you should call this option to establish an initial message filter.
// 
// An empty option_value of length zero shall subscribe to all incoming
// messages. A non-empty option_value shall subscribe to all messages beginning
// with the specified prefix. Multiple filters may be attached to a single
// ZMQ_SUB socket, in which case a message shall be accepted if it matches at
// least one filter.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             N/A
//  Applicable socket types   ZMQ_SUB
//
func (s *Socket) Subscribe(value []byte) error {
	return s.base.SetSockOptString(zmq.SUBSCRIBE, string(value))
}

// ZMQ_UNSUBSCRIBE: Remove message filter.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_UNSUBSCRIBE option shall remove an existing message filter on a
// ZMQ_SUB socket. The filter specified must match an existing filter previously
// established with the ZMQ_SUBSCRIBE option. If the socket has several
// instances of the same filter attached the ZMQ_UNSUBSCRIBE option shall remove
// only one instance, leaving the rest in place and functional.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             N/A
//  Applicable socket types   ZMQ_SUB
//
func (s *Socket) Unsubscribe(value []byte) error {
	return s.base.SetSockOptString(zmq.UNSUBSCRIBE, string(value))
}

// ZMQ_IDENTITY: Set socket identity.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_IDENTITY option shall set the identity of the specified socket.
// Socket identity is used only by request/reply pattern. Namely, it can be used
// in tandem with ROUTER socket to route messages to the peer with specific
// identity.
// 
// Identity should be at least one byte and at most 255 bytes long. Identities
// starting with binary zero are reserved for use by ØMQ infrastructure.
// 
// If two peers use the same identity when connecting to a third peer, the
// results shall be undefined.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             NULL
//  Applicable socket types   all
//
func (s *Socket) SetIdentity(value []byte) error {
	return s.base.SetSockOptString(zmq.IDENTITY, string(value))
}

// ZMQ_RATE: Set multicast data rate.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_RATE option shall set the maximum send or receive data rate for
// multicast transports such as zmq_pgm(7) using the specified socket.
// 
//  Option value type         int
//  Option value unit         kilobits per second
//  Default value             100
//  Applicable socket types   all, when using multicast transports
//
func (s *Socket) SetRate(value int64) error {
	return s.base.SetSockOptInt64(zmq.RATE, value)
}

// ZMQ_RECOVERY_IVL: Set multicast recovery interval.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_RECOVERY_IVL option shall set the recovery interval for multicast
// transports using the specified socket. The recovery interval determines the
// maximum time in milliseconds that a receiver can be absent from a multicast
// group before unrecoverable data loss will occur.
// 
// Exercise care when setting large recovery intervals as the data needed for
// recovery will be held in memory. For example, a 1 minute recovery interval at
// a data rate of 1Gbps requires a 7GB in-memory buffer.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             10000
//  Applicable socket types   all, when using multicast transports
//
func (s *Socket) SetRecoveryIvl(value time.Duration) error {
	return s.base.SetSockOptInt64(zmq.RECOVERY_IVL, int64(value/time.Millisecond))
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_SNDBUF option shall set the underlying kernel transmit buffer size
// for the socket to the specified size in bytes. A value of zero means leave
// the OS default unchanged. For details please refer to your operating system
// documentation for the SO_SNDBUF socket option.
// 
//  Option value type         int
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
func (s *Socket) SetSndbuf(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.SNDBUF, value)
}

// ZMQ_RCVBUF: Set kernel receive buffer size.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_RCVBUF option shall set the underlying kernel receive buffer size
// for the socket to the specified size in bytes. A value of zero means leave
// the OS default unchanged. For details refer to your operating system
// documentation for the SO_RCVBUF socket option.
// 
//  Option value type         int
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
func (s *Socket) SetRcvbuf(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.RCVBUF, value)
}

// ZMQ_LINGER: Set linger period for socket shutdown.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_LINGER option shall set the linger period for the specified socket.
// The linger period determines how long pending messages which have yet to be
// sent to a peer shall linger in memory after a socket is closed with
// zmq_close(3), and further affects the termination of the socket's context
// with zmq_term(3). The following outlines the different behaviours:
// 
// The default value of -1 specifies an infinite linger period. Pending
// messages shall not be discarded after a call to zmq_close(); attempting to
// terminate the socket's context with zmq_term() shall block until all pending
// messages have been sent to a peer.
// 
// The value of 0 specifies no linger period. Pending messages shall be
// discarded immediately when the socket is closed with zmq_close().
// 
// Positive values specify an upper bound for the linger period in
// milliseconds. Pending messages shall not be discarded after a call to
// zmq_close(); attempting to terminate the socket's context with zmq_term()
// shall block until either all pending messages have been sent to a peer, or
// the linger period expires, after which any pending messages shall be
// discarded.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.LINGER, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_RECONNECT_IVL option shall set the initial reconnection interval for
// the specified socket. The reconnection interval is the period ØMQ shall wait
// between attempts to reconnect disconnected peers when using connection-
// oriented transports. The value -1 means no reconnection.
// 
// The reconnection interval may be randomized by ØMQ to prevent reconnection
// storms in topologies with a large number of peers per socket.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             100
//  Applicable socket types   all, only for connection-oriented transports
//
func (s *Socket) SetReconnectIvl(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RECONNECT_IVL, int(value/time.Millisecond))
}

// ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_RECONNECT_IVL_MAX option shall set the maximum reconnection interval
// for the specified socket. This is the maximum period ØMQ shall wait between
// attempts to reconnect. On each reconnect attempt, the previous interval shall
// be doubled untill ZMQ_RECONNECT_IVL_MAX is reached. This allows for
// exponential backoff strategy. Default value means no exponential backoff is
// performed and reconnect interval calculations are only based on
// ZMQ_RECONNECT_IVL.
// 
// Values less than ZMQ_RECONNECT_IVL will be ignored.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             0 (only use ZMQ_RECONNECT_IVL)
//  Applicable socket types   all, only for connection-oriented transports
//
func (s *Socket) SetReconnectIvlMax(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RECONNECT_IVL_MAX, int(value/time.Millisecond))
}

// ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// The ZMQ_BACKLOG option shall set the maximum length of the queue of
// outstanding peer connections for the specified socket; this only applies to
// connection-oriented transports. For details refer to your operating system
// documentation for the listen function.
// 
//  Option value type         int
//  Option value unit         connections
//  Default value             100
//  Applicable socket types   all, only for connection-oriented transports.
//
func (s *Socket) SetBacklog(value int) error {
	return s.base.SetSockOptInt(zmq.BACKLOG, value)
}

// ZMQ_MAXMSGSIZE: Maximum acceptable inbound message size.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Limits the size of the inbound message. If a peer sends a message larger
// than ZMQ_MAXMSGSIZE it is disconnected. Value of -1 means no limit.
// 
//  Option value type         int64_t
//  Option value unit         bytes
//  Default value             -1
//  Applicable socket types   all
//
func (s *Socket) SetMaxMsgSize(value int64) error {
	return s.base.SetSockOptInt64(zmq.MAXMSGSIZE, value)
}

// ZMQ_MULTICAST_HOPS: Maximum network hops for multicast packets.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Sets the time-to-live field in every multicast packet sent from this socket.
// The default is 1 which means that the multicast packets don't leave the local
// network.
// 
//  Option value type         int
//  Option value unit         network hops
//  Default value             1
//  Applicable socket types   all, when using multicast transports
//
func (s *Socket) SetMulticastHops(value int) error {
	return s.base.SetSockOptInt(zmq.MULTICAST_HOPS, value)
}

// ZMQ_RCVTIMEO: Maximum time before a recv operation returns with EAGAIN.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Sets the timeout for receive operation on the socket. If the value is 0,
// zmq_recv(3) will return immediately, with a EAGAIN error if there is no
// message to receive. If the value is -1, it will block until a message is
// available. For all other values, it will wait for a message for that amount
// of time before returning with an EAGAIN error.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
func (s *Socket) SetRcvTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RCVTIMEO, int(value/time.Millisecond))
}

// ZMQ_SNDTIMEO: Maximum time before a send operation returns with EAGAIN.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Sets the timeout for send operation on the socket. If the value is 0,
// zmq_send(3) will return immediately, with a EAGAIN error if the message
// cannot be sent. If the value is -1, it will block until the message is sent.
// For all other values, it will try to send the message for that amount of time
// before returning with an EAGAIN error.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
func (s *Socket) SetSndTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.SNDTIMEO, int(value/time.Millisecond))
}

// ZMQ_IPV4ONLY: Use IPv4-only sockets.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Sets the underlying native socket type. A value of 1 will use IPv4 sockets,
// while the value of 0 will use IPv6 sockets. An IPv6 socket lets applications
// connect to and accept connections from both IPv4 and IPv6 hosts.
// 
//  Option value type         int
//  Option value unit         boolean
//  Default value             1 (true)
//  Applicable socket types   all, when using TCP transports.
//
func (s *Socket) SetIPV4Only(value int) error {
	return s.base.SetSockOptInt(zmq.IPV4ONLY, value)
}

// ZMQ_DELAY_ATTACH_ON_CONNECT: Accept messages only when connections are made.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// If set to 1, will delay the attachment of a pipe on connect until the
// underlying connection has completed. This will cause the socket to block if
// there are no other connections, but will prevent queues from filling on pipes
// awaiting connection.
// 
//  Option value type         int
//  Option value unit         boolean
//  Default value             0 (false)
//  Applicable socket types   all, only for connection-oriented transports.
//
func (s *Socket) SetDelayAttachOnConnect(value int) error {
	return s.base.SetSockOptInt(zmq.DELAY_ATTACH_ON_CONNECT, value)
}

// ZMQ_ROUTER_MANDATORY: accept only routable messages on ROUTER sockets.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Sets the ROUTER socket behavior when an unroutable message is encountered. A
// value of 0 is the default and discards the message silently when it cannot be
// routed. A value of 1 returns an EHOSTUNREACH error code if the message cannot
// be routed.
// 
//  Option value type         int
//  Option value unit         0, 1
//  Default value             0
//  Applicable socket types   ZMQ_ROUTER
//
func (s *Socket) SetRouterMandatory(value int) error {
	return s.base.SetSockOptInt(zmq.ROUTER_MANDATORY, value)
}

// ZMQ_XPUB_VERBOSE: provide all subscription messages on XPUB sockets.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Sets the XPUB socket behavior on new subscriptions and unsubscriptions. A
// value of 0 is the default and passes only new subscription messages to
// upstream. A value of 1 passes all subscription messages upstream.
// 
//  Option value type         int
//  Option value unit         0, 1
//  Default value             0
//  Applicable socket types   ZMQ_XPUB
//
func (s *Socket) SetXPUBVerbose(value int) error {
	return s.base.SetSockOptInt(zmq.XPUB_VERBOSE, value)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Override SO_KEEPALIVE socket option(where supported by OS). The default
// value of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0,1
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
func (s *Socket) SetTCPKeepalive(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE, value)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS) socket option(where
// supported by OS). The default value of -1 means to skip any overrides and
// leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
func (s *Socket) SetTCPKeepaliveIdle(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_IDLE, value)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Override TCP_KEEPCNT socket option(where supported by OS). The default value
// of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
func (s *Socket) SetTCPKeepaliveCnt(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_CNT, value)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Override TCP_KEEPINTVL socket option(where supported by OS). The default
// value of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
func (s *Socket) SetTCPKeepaliveIntvl(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_INTVL, value)
}

// ZMQ_TCP_ACCEPT_FILTER: Assign filters to allow new TCP connections.
// 
// From http://api.zeromq.org/3-2:zmq-setsockopt:
// 
// Assign arbitrary number of filters that will be applied for each new TCP
// transport connection on a listening socket. If no filters applied, then TCP
// transport allows connections from any ip. If at least one filter is applied
// then new connection source ip should be matched. To clear all filters call
// zmq_setsockopt(socket, ZMQ_TCP_ACCEPT_FILTER, NULL, 0). Filter is a null-
// terminated string with ipv6 or ipv4 CIDR.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             no filters (allow from all)
//  Applicable socket types   all listening sockets, when using TCP transports.
//
func (s *Socket) SetTCPAcceptFilter(value []byte) error {
	return s.base.SetSockOptString(zmq.TCP_ACCEPT_FILTER, string(value))
}
