// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package zmqutil implements some ØMQ abstractions and utilities.
package zmqutil

type _err int

const (
	_                 = iota
	ContextIsNil _err = _err(iota)
	SocketIsNil
	NotImplemented
)

func (e _err) Error() string {
	switch e {
	case ContextIsNil:
		return "gctx: nil Context"
	case SocketIsNil:
		return "gctx: nil Socket"
	case NotImplemented:
		return "gctx: this function is not yet implemented."
	}
	return "unknown error"
}
