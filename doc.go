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

/*
Package zmqutil implements some ØMQ (http://www.zeromq.org) abstractions
and utilities.

A context from this package remembers its sockets and has its own Linger
option.  When a context is closed, it will set the Linger option on each
socket that would linger longer and then close them all.

An additonal type, Poller, provides a convenient way to attach event
handlers to sockets.
*/
package zmqutil
