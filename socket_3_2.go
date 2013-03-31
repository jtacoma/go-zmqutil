// +build zmq_3_x
//
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
	"time"

	zmq "github.com/alecthomas/gozmq"
)

// This file was generated automatically.  Changes made here will be lost.

// Socket Option Getters

// ZMQ_TYPE: Retrieve socket type.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc3
//
func (s *Socket) Type() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.TYPE)
}

// ZMQ_RCVMORE: More message data parts to follow.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc4
//
func (s *Socket) Rcvmore() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.RCVMORE)
}

// ZMQ_SNDHWM: Retrieves high water mark for outbound messages.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc5
//
func (s *Socket) SndHWM() (int, error) {
	return s.base.GetSockOptInt(zmq.SNDHWM)
}

// ZMQ_RCVHWM: Retrieve high water mark for inbound messages.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc6
//
func (s *Socket) RcvHWM() (int, error) {
	return s.base.GetSockOptInt(zmq.RCVHWM)
}

// ZMQ_AFFINITY: Retrieve I/O thread affinity.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc7
//
func (s *Socket) Affinity() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.AFFINITY)
}

// ZMQ_IDENTITY: Set socket identity.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc8
//
func (s *Socket) Identity() (string, error) {
	return s.base.GetSockOptString(zmq.IDENTITY)
}

// ZMQ_RATE: Retrieve multicast data rate.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc9
//
func (s *Socket) Rate() (int64, error) {
	return s.base.GetSockOptInt64(zmq.RATE)
}

// ZMQ_RECOVERY_IVL: Get multicast recovery interval.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc10
//
func (s *Socket) RecoveryIvl() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt64(zmq.RECOVERY_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDBUF: Retrieve kernel transmit buffer size.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc11
//
func (s *Socket) Sndbuf() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.SNDBUF)
}

// ZMQ_RCVBUF: Retrieve kernel receive buffer size.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc12
//
func (s *Socket) Rcvbuf() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.RCVBUF)
}

// ZMQ_LINGER: Retrieve linger period for socket shutdown.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc13
//
func (s *Socket) Linger() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.LINGER)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL: Retrieve reconnection interval.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc14
//
func (s *Socket) ReconnectIvl() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.RECONNECT_IVL)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_RECONNECT_IVL_MAX: Retrieve maximum reconnection interval.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc15
//
func (s *Socket) ReconnectIvlMax() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.RECONNECT_IVL_MAX)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_BACKLOG: Retrieve maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc16
//
func (s *Socket) Backlog() (int, error) {
	return s.base.GetSockOptInt(zmq.BACKLOG)
}

// ZMQ_RCVTIMEO: Maximum time before a socket operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc19
//
func (s *Socket) RcvTimeout() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.RCVTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_SNDTIMEO: Maximum time before a socket operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc20
//
func (s *Socket) SndTimeout() (time.Duration, error) {
	ms, err := s.base.GetSockOptInt(zmq.SNDTIMEO)
	return time.Duration(ms) * time.Millisecond, err
}

// ZMQ_EVENTS: Retrieve socket event state.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc24
//
func (s *Socket) Events() (uint64, error) {
	return s.base.GetSockOptUInt64(zmq.EVENTS)
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc26
//
func (s *Socket) TCPKeepalive() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc27
//
func (s *Socket) TCPKeepaliveIdle() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE_IDLE)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc28
//
func (s *Socket) TCPKeepaliveCnt() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE_CNT)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
//
// See: http://api.zeromq.org/3-2:zmq-getsockopt#toc29
//
func (s *Socket) TCPKeepaliveIntvl() (int, error) {
	return s.base.GetSockOptInt(zmq.TCP_KEEPALIVE_INTVL)
}

// Socket Option Setters

// ZMQ_SNDHWM: Set high water mark for outbound messages.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc3
//
func (s *Socket) SetSndHWM(value int) error {
	return s.base.SetSockOptInt(zmq.SNDHWM, value)
}

// ZMQ_RCVHWM: Set high water mark for inbound messages.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc4
//
func (s *Socket) SetRcvHWM(value int) error {
	return s.base.SetSockOptInt(zmq.RCVHWM, value)
}

// ZMQ_AFFINITY: Set I/O thread affinity.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc5
//
func (s *Socket) SetAffinity(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.AFFINITY, value)
}

// ZMQ_SUBSCRIBE: Establish message filter.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc6
//
func (s *Socket) SetSubscribe(value string) error {
	return s.base.SetSockOptString(zmq.SUBSCRIBE, value)
}

// ZMQ_UNSUBSCRIBE: Remove message filter.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc7
//
func (s *Socket) SetUnsubscribe(value string) error {
	return s.base.SetSockOptString(zmq.UNSUBSCRIBE, value)
}

// ZMQ_IDENTITY: Set socket identity.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc8
//
func (s *Socket) SetIdentity(value string) error {
	return s.base.SetSockOptString(zmq.IDENTITY, value)
}

// ZMQ_RATE: Set multicast data rate.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc9
//
func (s *Socket) SetRate(value int64) error {
	return s.base.SetSockOptInt64(zmq.RATE, value)
}

// ZMQ_RECOVERY_IVL: Set multicast recovery interval.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc10
//
func (s *Socket) SetRecoveryIvl(value time.Duration) error {
	return s.base.SetSockOptInt64(zmq.RECOVERY_IVL, int64(time.Duration(value)/time.Millisecond))
}

// ZMQ_SNDBUF: Set kernel transmit buffer size.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc11
//
func (s *Socket) SetSndbuf(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.SNDBUF, value)
}

// ZMQ_RCVBUF: Set kernel receive buffer size.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc12
//
func (s *Socket) SetRcvbuf(value uint64) error {
	return s.base.SetSockOptUInt64(zmq.RCVBUF, value)
}

// ZMQ_LINGER: Set linger period for socket shutdown.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc13
//
func (s *Socket) SetLinger(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.LINGER, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_RECONNECT_IVL: Set reconnection interval.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc14
//
func (s *Socket) SetReconnectIvl(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RECONNECT_IVL, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_RECONNECT_IVL_MAX: Set maximum reconnection interval.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc15
//
func (s *Socket) SetReconnectIvlMax(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RECONNECT_IVL_MAX, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_BACKLOG: Set maximum length of the queue of outstanding connections.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc16
//
func (s *Socket) SetBacklog(value int) error {
	return s.base.SetSockOptInt(zmq.BACKLOG, value)
}

// ZMQ_RCVTIMEO: Maximum time before a recv operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc19
//
func (s *Socket) SetRcvTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.RCVTIMEO, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_SNDTIMEO: Maximum time before a send operation returns with EAGAIN.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc20
//
func (s *Socket) SetSndTimeout(value time.Duration) error {
	return s.base.SetSockOptInt(zmq.SNDTIMEO, int(time.Duration(value)/time.Millisecond))
}

// ZMQ_TCP_KEEPALIVE: Override SO_KEEPALIVE socket option.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc25
//
func (s *Socket) SetTCPKeepalive(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE, value)
}

// ZMQ_TCP_KEEPALIVE_IDLE: Override TCP_KEEPCNT(or TCP_KEEPALIVE on some OS).
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc26
//
func (s *Socket) SetTCPKeepaliveIdle(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_IDLE, value)
}

// ZMQ_TCP_KEEPALIVE_CNT: Override TCP_KEEPCNT socket option.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc27
//
func (s *Socket) SetTCPKeepaliveCnt(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_CNT, value)
}

// ZMQ_TCP_KEEPALIVE_INTVL: Override TCP_KEEPINTVL socket option.
//
// See: http://api.zeromq.org/3-2:zmq-setsockopt#toc28
//
func (s *Socket) SetTCPKeepaliveIntvl(value int) error {
	return s.base.SetSockOptInt(zmq.TCP_KEEPALIVE_INTVL, value)
}
