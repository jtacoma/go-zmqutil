# gzmq

[ØMQ](http://www.zeromq.org) abstractions for [Go](http://golang.org).

This package aims to implement, first, the [recommended binding abstractions](http://www.zeromq.org/topics:binding-abstractions).

## Status

This project is in early development, the public API may change significantly before 1.0.

## Implemented Features

* A context remembers its sockets and, when the context is closed, it will close all its sockets too.
* Linger is a context option, applied to its sockets when the context is closed.
* Reactor loop for socket events.
* All options are available through option-specific getter/setter methods (in progress: not all options supported).

## Planned Features

* Bind method that allows "bind to random free port" (and returns resulting port).
* The build-tag buck stops here: unsupported or obsolete methods return errors at runtime.

## License

Use of this source code is governed by a BSD-style license that can be found in
the LICENSE file.
