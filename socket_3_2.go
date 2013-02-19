// +build zmq_3_x
//

package gozmqutil

import (
	"time"

	zmq "github.com/alecthomas/gozmq"
)

// This file was generated automatically.  Changes made here will be lost.

// Socket Option Getters

// ZMQ_TYPE: Retrieve socket type.
// 
// The ZMQ_TYPE option shall retrieve the socket type for the specified
// socket. The socket type is specified at socket creation time and cannot
// be modified afterwards.
// 
//  Option value type         int
//  Option value unit         N/A
//  Default value             N/A
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc3)
//
func (s *Socket) Type() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.TYPE)
}

// ZMQ_RCVMORE: More message data parts to follow.
// 
// The ZMQ_RCVMORE option shall return True (1) if the message part last
// received from the socket was a data part with more parts to follow. If
// there are no data parts to follow, this option shall return False (0).
// 
// Refer to zmq_send(3) and zmq_recv(3) for a detailed description of
// multi-part messages.
// 
//  Option value type         int
//  Option value unit         boolean
//  Default value             N/A
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc4)
//
func (s *Socket) Rcvmore() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.RCVMORE)
}

// ZMQ_SNDHWM: Retrieves high water mark for outbound messages.
// 
// The ZMQ_SNDHWM option shall return the high water mark for outbound
// messages on the specified socket. The high water mark is a hard limit on
// the maximum number of outstanding messages ØMQ shall queue in memory for
// any single peer that the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional
// state and depending on the socket type, ØMQ shall take appropriate
// action such as blocking or dropping sent messages. Refer to the
// individual socket descriptions in zmq_socket(3) for details on the exact
// action taken for each socket type.
// 
//  Option value type         int
//  Option value unit         messages
//  Default value             1000
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc5)
//
func (s *Socket) SndHWM() (int, error) {
	return s.base.GetSockOptInt(zmq.SNDHWM)
}

// ZMQ_RCVHWM: Retrieve high water mark for inbound messages.
// 
// The ZMQ_RCVHWM option shall return the high water mark for inbound
// messages on the specified socket. The high water mark is a hard limit on
// the maximum number of outstanding messages ØMQ shall queue in memory for
// any single peer that the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional
// state and depending on the socket type, ØMQ shall take appropriate
// action such as blocking or dropping sent messages. Refer to the
// individual socket descriptions in zmq_socket(3) for details on the exact
// action taken for each socket type.
// 
//  Option value type         int
//  Option value unit         messages
//  Default value             1000
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc6)
//
func (s *Socket) RcvHWM() (int, error) {
	return s.base.GetSockOptInt(zmq.RCVHWM)
}

// ZMQ_AFFINITY: Retrieve I/O thread affinity.
// 
// The ZMQ_AFFINITY option shall retrieve the I/O thread affinity for
// newly created connections on the specified socket.
// 
// Affinity determines which threads from the ØMQ I/O thread pool
// associated with the socket's context shall handle newly created
// connections. A value of zero specifies no affinity, meaning that work
// shall be distributed fairly among all ØMQ I/O threads in the thread
// pool. For non-zero values, the lowest bit corresponds to thread 1,
// second lowest bit to thread 2 and so on. For example, a value of 3
// specifies that subsequent connections on socket shall be handled
// exclusively by I/O threads 1 and 2.
// 
// See also zmq_init(3) for details on allocating the number of I/O
// threads for a specific context.
// 
//  Option value type         uint64_t
//  Option value unit         N/A (bitmap)
//  Default value             0
//  Applicable socket types   N/A
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc7)
//
func (s *Socket) Affinity() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.AFFINITY)
}

// ZMQ_IDENTITY: Set socket identity.
// 
// The ZMQ_IDENTITY option shall retrieve the identity of the specified
// socket. Socket identity is used only by request/reply pattern. Namely,
// it can be used in tandem with ROUTER socket to route messages to the
// peer with specific identity.
// 
// Identity should be at least one byte and at most 255 bytes long.
// Identities starting with binary zero are reserved for use by ØMQ
// infrastructure.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             NULL
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc8)
//
func (s *Socket) Identity() (string, error) {
	return s.base.GetSockOptString(zmq.IDENTITY)
}

// ZMQ_RATE: Retrieve multicast data rate.
// 
// The ZMQ_RATE option shall retrieve the maximum send or receive data
// rate for multicast transports using the specified socket.
// 
//  Option value type         int
//  Option value unit         kilobits per second
//  Default value             100
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc9)
//
func (s *Socket) Rate() (int64, error) {
	return s.base.GetSockOptInt64(zmq.RATE)
}

// ZMQ_RECOVERY_IVL: Get multicast recovery interval.
// 
// The ZMQ_RECOVERY_IVL option shall retrieve the recovery interval for
// multicast transports using the specified socket. The recovery interval
// determines the maximum time in milliseconds that a receiver can be
// absent from a multicast group before unrecoverable data loss will occur.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             10000
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc10)
//
func (s *Socket) RecoveryIvl() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt64(zmq.RECOVERY_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDBUF: Retrieve kernel transmit buffer size.
// 
// The ZMQ_SNDBUF option shall retrieve the underlying kernel transmit
// buffer size for the specified socket. A value of zero means that the OS
// default is in effect. For details refer to your operating system
// documentation for the SO_SNDBUF socket option.
// 
//  Option value type         int
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc11)
//
func (s *Socket) Sndbuf() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.SNDBUF)
}

// ZMQ_RCVBUF: Retrieve kernel receive buffer size.
// 
// The ZMQ_RCVBUF option shall retrieve the underlying kernel receive
// buffer size for the specified socket. A value of zero means that the OS
// default is in effect. For details refer to your operating system
// documentation for the SO_RCVBUF socket option.
// 
//  Option value type         int
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc12)
//
func (s *Socket) Rcvbuf() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.RCVBUF)
}

// ZMQ_LINGER: Retrieve linger period for socket shutdown.
// 
// The ZMQ_LINGER option shall retrieve the linger period for the
// specified socket. The linger period determines how long pending messages
// which have yet to be sent to a peer shall linger in memory after a
// socket is closed with zmq_close(3), and further affects the termination
// of the socket's context with zmq_term(3). The following outlines the
// different behaviours:
// 
// The default value of -1 specifies an infinite linger period. Pending
// messages shall not be discarded after a call to zmq_close(); attempting
// to terminate the socket's context with zmq_term() shall block until all
// pending messages have been sent to a peer.
// 
// The value of 0 specifies no linger period. Pending messages shall be
// discarded immediately when the socket is closed with zmq_close().
// 
// Positive values specify an upper bound for the linger period in
// milliseconds. Pending messages shall not be discarded after a call to
// zmq_close(); attempting to terminate the socket's context with
// zmq_term() shall block until either all pending messages have been sent
// to a peer, or the linger period expires, after which any pending
// messages shall be discarded.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc13)
//
func (s *Socket) Linger() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.LINGER)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL: Retrieve reconnection interval.
// 
// The ZMQ_RECONNECT_IVL option shall retrieve the initial reconnection
// interval for the specified socket. The reconnection interval is the
// period ØMQ shall wait between attempts to reconnect disconnected peers
// when using connection-oriented transports. The value -1 means no
// reconnection.
// 
// The reconnection interval may be randomized by ØMQ to prevent
// reconnection storms in topologies with a large number of peers per
// socket.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             100
//  Applicable socket types   all, only for connection-oriented transports
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc14)
//
func (s *Socket) ReconnectIvl() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.RECONNECT_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL_MAX: Retrieve maximum reconnection interval.
// 
// The ZMQ_RECONNECT_IVL_MAX option shall retrieve the maximum
// reconnection interval for the specified socket. This is the maximum
// period ØMQ shall wait between attempts to reconnect. On each reconnect
// attempt, the previous interval shall be doubled untill
// ZMQ_RECONNECT_IVL_MAX is reached. This allows for exponential backoff
// strategy. Default value means no exponential backoff is performed and
// reconnect interval calculations are only based on ZMQ_RECONNECT_IVL.
// 
// Values less than ZMQ_RECONNECT_IVL will be ignored.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             0 (only use ZMQ_RECONNECT_IVL)
//  Applicable socket types   all, only for connection-oriented transport
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc15)
//
func (s *Socket) ReconnectIvlMax() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.RECONNECT_IVL_MAX)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_BACKLOG: Retrieve maximum length of the queue of outstanding
// connections.
// 
// The ZMQ_BACKLOG option shall retrieve the maximum length of the queue
// of outstanding peer connections for the specified socket; this only
// applies to connection-oriented transports. For details refer to your
// operating system documentation for the listen function.
// 
//  Option value type         int
//  Option value unit         connections
//  Default value             100
//  Applicable socket types   all, only for connection-oriented transports
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc16)
//
func (s *Socket) Backlog() (int, error) {
	return s.base.GetSockOptInt(zmq.BACKLOG)
}

// ZMQ_RCVTIMEO: Maximum time before a socket operation returns with
// EAGAIN.
// 
// Retrieve the timeout for recv operation on the socket. If the value is
// 0, zmq_recv(3) will return immediately, with a EAGAIN error if there is
// no message to receive. If the value is -1, it will block until a message
// is available. For all other values, it will wait for a message for that
// amount of time before returning with an EAGAIN error.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc19)
//
func (s *Socket) RcvTimeout() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.RCVTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDTIMEO: Maximum time before a socket operation returns with
// EAGAIN.
// 
// Retrieve the timeout for send operation on the socket. If the value is
// 0, zmq_send(3) will return immediately, with a EAGAIN error if the
// message cannot be sent. If the value is -1, it will block until the
// message is sent. For all other values, it will try to send the message
// for that amount of time before returning with an EAGAIN error.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc20)
//
func (s *Socket) SndTimeout() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.SNDTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_EVENTS: Retrieve socket event state.
// 
// The ZMQ_EVENTS option shall retrieve the event state for the specified
// socket. The returned value is a bit mask constructed by OR'ing a
// combination of the following event flags:
// 
// ZMQ_POLLIN Indicates that at least one message may be received from
// the specified socket without blocking. ZMQ_POLLOUT Indicates that at
// least one message may be sent to the specified socket without blocking.
// The combination of a file descriptor returned by the ZMQ_FD option being
// ready for reading but no actual events returned by a subsequent
// retrieval of the ZMQ_EVENTS option is valid; applications should simply
// ignore this case and restart their polling operation/event loop.
// 
//  Option value type         int
//  Option value unit         N/A (flags)
//  Default value             N/A
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc24)
//
func (s *Socket) Events() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.EVENTS)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
// 
// Override SO_KEEPALIVE socket option(where supported by OS). The default
// value of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0,1
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc26)
//
func (s *Socket) TCPKeepalive() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some
// OS).
// 
// Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS) socket option(where
// supported by OS). The default value of -1 means to skip any overrides
// and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc27)
//
func (s *Socket) TCPKeepaliveIdle() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE_IDLE)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
// 
// Override TCP_KEEPCNT socket option(where supported by OS). The default
// value of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc28)
//
func (s *Socket) TCPKeepaliveCnt() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE_CNT)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
// 
// Override TCP_KEEPINTVL socket option(where supported by OS). The
// default value of -1 means to skip any overrides and leave it to OS
// default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-getsockopt#toc29)
//
func (s *Socket) TCPKeepaliveIntvl() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE_INTVL)
}

// Socket Option Setters

// ZMQ_SNDHWM: Set high water mark for outbound messages.
// 
// The ZMQ_SNDHWM option shall set the high water mark for outbound
// messages on the specified socket. The high water mark is a hard limit on
// the maximum number of outstanding messages ØMQ shall queue in memory for
// any single peer that the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional
// state and depending on the socket type, ØMQ shall take appropriate
// action such as blocking or dropping sent messages. Refer to the
// individual socket descriptions in zmq_socket(3) for details on the exact
// action taken for each socket type.
// 
// ØMQ does not guarantee that the socket will accept as many as
// ZMQ_SNDHWM messages, and the actual limit may be as much as 60-70% lower
// depending on the flow of messages on the socket.
// 
//  Option value type         int
//  Option value unit         messages
//  Default value             1000
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc3)
//
func (s *Socket) SetSndHWM(value int) error {
	return s.base.SetSockOptInt(zmq.SNDHWM, value)
}

// ZMQ_RCVHWM: Set high water mark for inbound messages.
// 
// The ZMQ_RCVHWM option shall set the high water mark for inbound
// messages on the specified socket. The high water mark is a hard limit on
// the maximum number of outstanding messages ØMQ shall queue in memory for
// any single peer that the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional
// state and depending on the socket type, ØMQ shall take appropriate
// action such as blocking or dropping sent messages. Refer to the
// individual socket descriptions in zmq_socket(3) for details on the exact
// action taken for each socket type.
// 
//  Option value type         int
//  Option value unit         messages
//  Default value             1000
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc4)
//
func (s *Socket) SetRcvHWM(value int) error {
	return s.base.SetSockOptInt(zmq.RCVHWM, value)
}

// ZMQ_AFFINITY: Set I/O thread affinity.
// 
// The ZMQ_AFFINITY option shall set the I/O thread affinity for newly
// created connections on the specified socket.
// 
// Affinity determines which threads from the ØMQ I/O thread pool
// associated with the socket's context shall handle newly created
// connections. A value of zero specifies no affinity, meaning that work
// shall be distributed fairly among all ØMQ I/O threads in the thread
// pool. For non-zero values, the lowest bit corresponds to thread 1,
// second lowest bit to thread 2 and so on. For example, a value of 3
// specifies that subsequent connections on socket shall be handled
// exclusively by I/O threads 1 and 2.
// 
// See also zmq_init(3) for details on allocating the number of I/O
// threads for a specific context.
// 
//  Option value type         uint64_t
//  Option value unit         N/A (bitmap)
//  Default value             0
//  Applicable socket types   N/A
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc5)
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.AFFINITY, value)
}

// ZMQ_SUBSCRIBE: Establish message filter.
// 
// The ZMQ_SUBSCRIBE option shall establish a new message filter on a
// ZMQ_SUB socket. Newly created ZMQ_SUB sockets shall filter out all
// incoming messages, therefore you should call this option to establish an
// initial message filter.
// 
// An empty option_value of length zero shall subscribe to all incoming
// messages. A non-empty option_value shall subscribe to all messages
// beginning with the specified prefix. Multiple filters may be attached to
// a single ZMQ_SUB socket, in which case a message shall be accepted if it
// matches at least one filter.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             N/A
//  Applicable socket types   ZMQ_SUB
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc6)
//
func (s *Socket) SetSubscribe(value string) error {
	return s.base.SetSockOptString(zmq.SUBSCRIBE, value)
}

// ZMQ_UNSUBSCRIBE: Remove message filter.
// 
// The ZMQ_UNSUBSCRIBE option shall remove an existing message filter on a
// ZMQ_SUB socket. The filter specified must match an existing filter
// previously established with the ZMQ_SUBSCRIBE option. If the socket has
// several instances of the same filter attached the ZMQ_UNSUBSCRIBE option
// shall remove only one instance, leaving the rest in place and
// functional.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             N/A
//  Applicable socket types   ZMQ_SUB
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc7)
//
func (s *Socket) SetUnsubscribe(value string) error {
	return s.base.SetSockOptString(zmq.UNSUBSCRIBE, value)
}

// ZMQ_IDENTITY: Set socket identity.
// 
// The ZMQ_IDENTITY option shall set the identity of the specified socket.
// Socket identity is used only by request/reply pattern. Namely, it can be
// used in tandem with ROUTER socket to route messages to the peer with
// specific identity.
// 
// Identity should be at least one byte and at most 255 bytes long.
// Identities starting with binary zero are reserved for use by ØMQ
// infrastructure.
// 
// If two peers use the same identity when connecting to a third peer, the
// results shall be undefined.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             NULL
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc8)
//
func (s *Socket) SetIdentity(value string) error {
	return s.base.SetSockOptString(zmq.IDENTITY, value)
}

// ZMQ_RATE: Set multicast data rate.
// 
// The ZMQ_RATE option shall set the maximum send or receive data rate for
// multicast transports such as zmq_pgm(7) using the specified socket.
// 
//  Option value type         int
//  Option value unit         kilobits per second
//  Default value             100
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc9)
//
func (s *Socket) SetRate(value int64) error {
	return s.base.SetSockOptInt64(zmq.RATE, value)
}

// ZMQ_RECOVERY_IVL: Set multicast recovery interval.
// 
// The ZMQ_RECOVERY_IVL option shall set the recovery interval for
// multicast transports using the specified socket. The recovery interval
// determines the maximum time in milliseconds that a receiver can be
// absent from a multicast group before unrecoverable data loss will occur.
// 
// Exercise care when setting large recovery intervals as the data needed
// for recovery will be held in memory. For example, a 1 minute recovery
// interval at a data rate of 1Gbps requires a 7GB in-memory buffer.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             10000
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc10)
//
func (s *Socket) SetRecoveryIvl(value time.Duration) error {
	return s.base.SetSockOptInt64(zmq.RECOVERY_IVL, int64(time.Duration(value)/time.Millisecond))
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
// 
// The ZMQ_SNDBUF option shall set the underlying kernel transmit buffer
// size for the socket to the specified size in bytes. A value of zero
// means leave the OS default unchanged. For details please refer to your
// operating system documentation for the SO_SNDBUF socket option.
// 
//  Option value type         int
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc11)
//
func (s *Socket) SetSndbuf(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.SNDBUF, value)
}

// ZMQ_RCVBUF: Set kernel receive buffer size.
// 
// The ZMQ_RCVBUF option shall set the underlying kernel receive buffer
// size for the socket to the specified size in bytes. A value of zero
// means leave the OS default unchanged. For details refer to your
// operating system documentation for the SO_RCVBUF socket option.
// 
//  Option value type         int
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc12)
//
func (s *Socket) SetRcvbuf(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.RCVBUF, value)
}

// ZMQ_LINGER: Set linger period for socket shutdown.
// 
// The ZMQ_LINGER option shall set the linger period for the specified
// socket. The linger period determines how long pending messages which
// have yet to be sent to a peer shall linger in memory after a socket is
// closed with zmq_close(3), and further affects the termination of the
// socket's context with zmq_term(3). The following outlines the different
// behaviours:
// 
// The default value of -1 specifies an infinite linger period. Pending
// messages shall not be discarded after a call to zmq_close(); attempting
// to terminate the socket's context with zmq_term() shall block until all
// pending messages have been sent to a peer.
// 
// The value of 0 specifies no linger period. Pending messages shall be
// discarded immediately when the socket is closed with zmq_close().
// 
// Positive values specify an upper bound for the linger period in
// milliseconds. Pending messages shall not be discarded after a call to
// zmq_close(); attempting to terminate the socket's context with
// zmq_term() shall block until either all pending messages have been sent
// to a peer, or the linger period expires, after which any pending
// messages shall be discarded.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc13)
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.LINGER, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
// 
// The ZMQ_RECONNECT_IVL option shall set the initial reconnection
// interval for the specified socket. The reconnection interval is the
// period ØMQ shall wait between attempts to reconnect disconnected peers
// when using connection-oriented transports. The value -1 means no
// reconnection.
// 
// The reconnection interval may be randomized by ØMQ to prevent
// reconnection storms in topologies with a large number of peers per
// socket.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             100
//  Applicable socket types   all, only for connection-oriented transports
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc14)
//
func (s *Socket) SetReconnectIvl(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RECONNECT_IVL, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.
// 
// The ZMQ_RECONNECT_IVL_MAX option shall set the maximum reconnection
// interval for the specified socket. This is the maximum period ØMQ shall
// wait between attempts to reconnect. On each reconnect attempt, the
// previous interval shall be doubled untill ZMQ_RECONNECT_IVL_MAX is
// reached. This allows for exponential backoff strategy. Default value
// means no exponential backoff is performed and reconnect interval
// calculations are only based on ZMQ_RECONNECT_IVL.
// 
// Values less than ZMQ_RECONNECT_IVL will be ignored.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             0 (only use ZMQ_RECONNECT_IVL)
//  Applicable socket types   all, only for connection-oriented transports
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc15)
//
func (s *Socket) SetReconnectIvlMax(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RECONNECT_IVL_MAX, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.
// 
// The ZMQ_BACKLOG option shall set the maximum length of the queue of
// outstanding peer connections for the specified socket; this only applies
// to connection-oriented transports. For details refer to your operating
// system documentation for the listen function.
// 
//  Option value type         int
//  Option value unit         connections
//  Default value             100
//  Applicable socket types   all, only for connection-oriented transports.
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc16)
//
func (s *Socket) SetBacklog(value int) error {
	return s.base.SetSockOptInt(zmq.BACKLOG, value)
}

// ZMQ_RCVTIMEO: Maximum time before a recv operation returns with EAGAIN.
// 
// Sets the timeout for receive operation on the socket. If the value is
// 0, zmq_recv(3) will return immediately, with a EAGAIN error if there is
// no message to receive. If the value is -1, it will block until a message
// is available. For all other values, it will wait for a message for that
// amount of time before returning with an EAGAIN error.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc19)
//
func (s *Socket) SetRcvTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RCVTIMEO, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_SNDTIMEO: Maximum time before a send operation returns with EAGAIN.
// 
// Sets the timeout for send operation on the socket. If the value is 0,
// zmq_send(3) will return immediately, with a EAGAIN error if the message
// cannot be sent. If the value is -1, it will block until the message is
// sent. For all other values, it will try to send the message for that
// amount of time before returning with an EAGAIN error.
// 
//  Option value type         int
//  Option value unit         milliseconds
//  Default value             -1 (infinite)
//  Applicable socket types   all
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc20)
//
func (s *Socket) SetSndTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.SNDTIMEO, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
// 
// Override SO_KEEPALIVE socket option(where supported by OS). The default
// value of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0,1
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc25)
//
func (s *Socket) SetTCPKeepalive(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE, value)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some
// OS).
// 
// Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS) socket option(where
// supported by OS). The default value of -1 means to skip any overrides
// and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc26)
//
func (s *Socket) SetTCPKeepaliveIdle(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_IDLE, value)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
// 
// Override TCP_KEEPCNT socket option(where supported by OS). The default
// value of -1 means to skip any overrides and leave it to OS default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc27)
//
func (s *Socket) SetTCPKeepaliveCnt(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_CNT, value)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
// 
// Override TCP_KEEPINTVL socket option(where supported by OS). The
// default value of -1 means to skip any overrides and leave it to OS
// default.
// 
//  Option value type         int
//  Option value unit         -1,0
//  Default value             -1 (leave to OS default)
//  Applicable socket types   all, when using TCP transports.
//
// (from http://api.zeromq.org/3-2:zmq-setsockopt#toc28)
//
func (s *Socket) SetTCPKeepaliveIntvl(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_INTVL, value)
}
