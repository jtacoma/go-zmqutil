// +build !zmq_2_1,!zmq_3_x
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc3)
//
func (s *Socket) Type() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.TYPE)
}

// ZMQ_RCVMORE: More message parts to follow.
// 
// The ZMQ_RCVMORE option shall return a boolean value indicating if the
// multi-part message currently being read from the specified socket has
// more message parts to follow. If there are no message parts to follow or
// if the message currently being read is not a multi-part message a value
// of zero shall be returned. Otherwise, a value of 1 shall be returned.
// 
// Refer to zmq_send(3) and zmq_recv(3) for a detailed description of
// sending/receiving multi-part messages.
// 
//  Option value type         int64_t
//  Option value unit         boolean
//  Default value             N/A
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc4)
//
func (s *Socket) Rcvmore() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.RCVMORE)
}

// ZMQ_HWM: Retrieve high water mark.
// 
// The ZMQ_HWM option shall retrieve the high water mark for the specified
// socket. The high water mark is a hard limit on the maximum number of
// outstanding messages ØMQ shall queue in memory for any single peer that
// the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional
// state and depending on the socket type, ØMQ shall take appropriate
// action such as blocking or dropping sent messages. Refer to the
// individual socket descriptions in zmq_socket(3) for details on the exact
// action taken for each socket type.
// 
// The default ZMQ_HWM value of zero means no limit.
// 
//  Option value type         uint64_t
//  Option value unit         messages
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc5)
//
func (s *Socket) Hwm() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.HWM)
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc6)
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc7)
//
func (s *Socket) SndTimeout() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.SNDTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SWAP: Retrieve disk offload size.
// 
// The ZMQ_SWAP option shall retrieve the disk offload (swap) size for the
// specified socket. A socket which has ZMQ_SWAP set to a non-zero value
// may exceed its high water mark; in this case outstanding messages shall
// be offloaded to storage on disk rather than held in memory.
// 
// The value of ZMQ_SWAP defines the maximum size of the swap space in
// bytes.
// 
//  Option value type         int64_t
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc8)
//
func (s *Socket) Swap() (int64, error) {
	return s.base.GetSockOptInt64(zmq.SWAP)
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc9)
//
func (s *Socket) Affinity() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.AFFINITY)
}

// ZMQ_IDENTITY: Retrieve socket identity.
// 
// The ZMQ_IDENTITY option shall retrieve the identity of the specified
// socket. Socket identity determines if existing ØMQ infrastructure
// (message queues, forwarding devices) shall be identified with a specific
// application and persist across multiple runs of the application.
// 
// If the socket has no identity, each run of an application is completely
// separate from other runs. However, with identity set the socket shall
// re-use any existing ØMQ infrastructure configured by the previous
// run(s). Thus the application may receive messages that were sent in the
// meantime, message queue limits shall be shared with previous run(s) and
// so on.
// 
// Identity can be at least one byte and at most 255 bytes long.
// Identities starting with binary zero are reserved for use by ØMQ
// infrastructure.
// 
//  Option value type         binary data
//  Option value unit         N/A
//  Default value             NULL
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc10)
//
func (s *Socket) Identity() (string, error) {
	return s.base.GetSockOptString(zmq.IDENTITY)
}

// ZMQ_RATE: Retrieve multicast data rate.
// 
// The ZMQ_RATE option shall retrieve the maximum send or receive data
// rate for multicast transports using the specified socket.
// 
//  Option value type         int64_t
//  Option value unit         kilobits per second
//  Default value             100
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc11)
//
func (s *Socket) Rate() (int64, error) {
	return s.base.GetSockOptInt64(zmq.RATE)
}

// ZMQ_RECOVERY_IVL: Get multicast recovery interval.
// 
// The ZMQ_RECOVERY_IVL option shall retrieve the recovery interval for
// multicast transports using the specified socket. The recovery interval
// determines the maximum time in seconds that a receiver can be absent
// from a multicast group before unrecoverable data loss will occur.
// 
//  Option value type         int64_t
//  Option value unit         seconds
//  Default value             10
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc12)
//
func (s *Socket) RecoveryIvl() (int64, error) {
	return s.base.GetSockOptInt64(zmq.RECOVERY_IVL)
}

// ZMQ_RECOVERY_IVL_MSEC: Get multicast recovery interval in milliseconds.
// 
// The ZMQ_RECOVERY_IVL'_MSEC option shall retrieve the recovery interval,
// in milliseconds, for multicast transports using the specified 'socket.
// The recovery interval determines the maximum time in seconds that a
// receiver can be absent from a multicast group before unrecoverable data
// loss will occur.
// 
// For backward compatibility, the default value of ZMQ_RECOVERY_IVL_MSEC
// is -1 indicating that the recovery interval should be obtained from the
// ZMQ_RECOVERY_IVL option. However, if the ZMQ_RECOVERY_IVL_MSEC value is
// not zero, then it will take precedence, and be used.
// 
//  Option value type         int64_t
//  Option value unit         milliseconds
//  Default value             -1
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc13)
//
func (s *Socket) RecoveryIvlMsec() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt64(zmq.RECOVERY_IVL_MSEC)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_MCAST_LOOP: Control multicast loop-back.
// 
// The ZMQ_MCAST_LOOP option controls whether data sent via multicast
// transports can also be received by the sending host via loop-back. A
// value of zero indicates that the loop-back functionality is disabled,
// while the default value of 1 indicates that the loop-back functionality
// is enabled. Leaving multicast loop-back enabled when it is not required
// can have a negative impact on performance. Where possible, disable
// ZMQ_MCAST_LOOP in production environments.
// 
//  Option value type         int64_t
//  Option value unit         boolean
//  Default value             1
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc14)
//
func (s *Socket) McastLoop() (int64, error) {
	return s.base.GetSockOptInt64(zmq.MCAST_LOOP)
}

// ZMQ_SNDBUF: Retrieve kernel transmit buffer size.
// 
// The ZMQ_SNDBUF option shall retrieve the underlying kernel transmit
// buffer size for the specified socket. A value of zero means that the OS
// default is in effect. For details refer to your operating system
// documentation for the SO_SNDBUF socket option.
// 
//  Option value type         uint64_t
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc15)
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
//  Option value type         uint64_t
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc16)
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc17)
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
// when using connection-oriented transports.
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc18)
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc19)
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
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc20)
//
func (s *Socket) Backlog() (int, error) {
	return s.base.GetSockOptInt(zmq.BACKLOG)
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
//  Option value type         uint32_t
//  Option value unit         N/A (flags)
//  Default value             N/A
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-getsockopt#toc22)
//
func (s *Socket) Events() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.EVENTS)
}

// Socket Option Setters

// ZMQ_HWM: Set high water mark.
// 
// The ZMQ_HWM option shall set the high water mark for the specified
// socket. The high water mark is a hard limit on the maximum number of
// outstanding messages ØMQ shall queue in memory for any single peer that
// the specified socket is communicating with.
// 
// If this limit has been reached the socket shall enter an exceptional
// state and depending on the socket type, ØMQ shall take appropriate
// action such as blocking or dropping sent messages. Refer to the
// individual socket descriptions in zmq_socket(3) for details on the exact
// action taken for each socket type.
// 
// The default ZMQ_HWM value of zero means no limit.
// 
//  Option value type         uint64_t
//  Option value unit         messages
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc3)
//
func (s *Socket) SetHwm(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.HWM, value)
}

// ZMQ_SWAP: Set disk offload size.
// 
// The ZMQ_SWAP option shall set the disk offload (swap) size for the
// specified socket. A socket which has ZMQ_SWAP set to a non-zero value
// may exceed its high water mark; in this case outstanding messages shall
// be offloaded to storage on disk rather than held in memory.
// 
// The value of ZMQ_SWAP defines the maximum size of the swap space in
// bytes.
// 
//  Option value type         int64_t
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc4)
//
func (s *Socket) SetSwap(value int64) error {
	return s.base.SetSockOptInt64(zmq.SWAP, value)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc5)
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.AFFINITY, value)
}

// ZMQ_IDENTITY: Set socket identity.
// 
// The ZMQ_IDENTITY option shall set the identity of the specified socket.
// Socket identity determines if existing ØMQ infrastructure (message
// queues, forwarding devices) shall be identified with a specific
// application and persist across multiple runs of the application.
// 
// If the socket has no identity, each run of an application is completely
// separate from other runs. However, with identity set the socket shall
// re-use any existing ØMQ infrastructure configured by the previous
// run(s). Thus the application may receive messages that were sent in the
// meantime, message queue limits shall be shared with previous run(s) and
// so on.
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc6)
//
func (s *Socket) SetIdentity(value string) error {
	return s.base.SetSockOptString(zmq.IDENTITY, value)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc7)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc8)
//
func (s *Socket) SetUnsubscribe(value string) error {
	return s.base.SetSockOptString(zmq.UNSUBSCRIBE, value)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc9)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc10)
//
func (s *Socket) SetSndTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.SNDTIMEO, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_RATE: Set multicast data rate.
// 
// The ZMQ_RATE option shall set the maximum send or receive data rate for
// multicast transports such as zmq_pgm(7) using the specified socket.
// 
//  Option value type         int64_t
//  Option value unit         kilobits per second
//  Default value             100
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc11)
//
func (s *Socket) SetRate(value int64) error {
	return s.base.SetSockOptInt64(zmq.RATE, value)
}

// ZMQ_RECOVERY_IVL: Set multicast recovery interval.
// 
// The ZMQ_RECOVERY_IVL option shall set the recovery interval for
// multicast transports using the specified socket. The recovery interval
// determines the maximum time in seconds that a receiver can be absent
// from a multicast group before unrecoverable data loss will occur.
// 
// Exercise care when setting large recovery intervals as the data needed
// for recovery will be held in memory. For example, a 1 minute recovery
// interval at a data rate of 1Gbps requires a 7GB in-memory buffer.
// 
//  Option value type         int64_t
//  Option value unit         seconds
//  Default value             10
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc12)
//
func (s *Socket) SetRecoveryIvl(value int64) error {
	return s.base.SetSockOptInt64(zmq.RECOVERY_IVL, value)
}

// ZMQ_RECOVERY_IVL_MSEC: Set multicast recovery interval in milliseconds.
// 
// The ZMQ_RECOVERY_IVL_MSEC option shall set the recovery interval,
// specified in milliseconds (ms) for multicast transports using the
// specified socket. The recovery interval determines the maximum time in
// milliseconds that a receiver can be absent from a multicast group before
// unrecoverable data loss will occur.
// 
// A non-zero value of the ZMQ_RECOVERY_IVL_MSEC option will take
// precedence over the ZMQ_RECOVERY_IVL option, but since the default for
// the ZMQ_RECOVERY_IVL_MSEC is -1, the default is to use the
// ZMQ_RECOVERY_IVL option value.
// 
// Exercise care when setting large recovery intervals as the data needed
// for recovery will be held in memory. For example, a 1 minute recovery
// interval at a data rate of 1Gbps requires a 7GB in-memory buffer.
// 
//  Option value type         int64_t
//  Option value unit         milliseconds
//  Default value             -1
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc13)
//
func (s *Socket) SetRecoveryIvlMsec(value time.Duration) error {
	return s.base.SetSockOptInt64(zmq.RECOVERY_IVL_MSEC, int64(time.Duration(value)/time.Millisecond))
}

// ZMQ_MCAST_LOOP: Control multicast loop-back.
// 
// The ZMQ_MCAST_LOOP option shall control whether data sent via multicast
// transports using the specified socket can also be received by the
// sending host via loop-back. A value of zero disables the loop-back
// functionality, while the default value of 1 enables the loop-back
// functionality. Leaving multicast loop-back enabled when it is not
// required can have a negative impact on performance. Where possible,
// disable ZMQ_MCAST_LOOP in production environments.
// 
//  Option value type         int64_t
//  Option value unit         boolean
//  Default value             1
//  Applicable socket types   all, when using multicast transports
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc14)
//
func (s *Socket) SetMcastLoop(value int64) error {
	return s.base.SetSockOptInt64(zmq.MCAST_LOOP, value)
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
// 
// The ZMQ_SNDBUF option shall set the underlying kernel transmit buffer
// size for the socket to the specified size in bytes. A value of zero
// means leave the OS default unchanged. For details please refer to your
// operating system documentation for the SO_SNDBUF socket option.
// 
//  Option value type         uint64_t
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc15)
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
//  Option value type         uint64_t
//  Option value unit         bytes
//  Default value             0
//  Applicable socket types   all
//
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc16)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc17)
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.LINGER, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
// 
// The ZMQ_RECONNECT_IVL option shall set the initial reconnection
// interval for the specified socket. The reconnection interval is the
// period ØMQ shall wait between attempts to reconnect disconnected peers
// when using connection-oriented transports.
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc18)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc19)
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
// (from http://api.zeromq.org/2-2:zmq-setsockopt#toc20)
//
func (s *Socket) SetBacklog(value int) error {
	return s.base.SetSockOptInt(zmq.BACKLOG, value)
}
