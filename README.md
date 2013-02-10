# gzmq

[ØMQ](http://www.zeromq.org) abstractions for [Go](http://golang.org).

This package aims to implement, first, the [recommended binding abstractions](http://www.zeromq.org/topics:binding-abstractions).

## Status

This project is in early development, the public API may change significantly before 1.0.

## Implemented Features

* A context remembers its sockets and, when the context is closed, it will close all its sockets too.
* `LINGER` is context option, applied to its sockets.
* Reactor loop for socket events.

## Planned Features

* API for all socket options (get/set as appropriate)
* Bind method that allows "bind to random free port" (and returns resulting port).

## License

Use of this source code is governed by a BSD-style license that can be found in
the LICENSE file.
