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
	zmq "github.com/alecthomas/gozmq"
)

type Socket struct {
	*zmq.Socket
	s *zmq.Socket // Really? ...yes, really.
}

func (s *Socket) MustBind(addr string) {
	if err := s.Bind(addr); err != nil {
		panic(err.Error())
	}
}
func (s *Socket) MustConnect(addr string) {
	if err := s.Connect(addr); err != nil {
		panic(err.Error())
	}
}
