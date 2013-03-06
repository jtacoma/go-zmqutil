// Copyright 2013 Joshua Tacoma. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
		return "zmqutil: nil Context"
	case SocketIsNil:
		return "zmqutil: nil Socket"
	case NotImplemented:
		return "zmqutil: this function is not yet implemented."
	}
	return "unknown error"
}
