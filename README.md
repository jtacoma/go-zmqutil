# gozmqutil

    import (
        zmq "github.com/alecthomas/gozmq"
        zmqutil "github.com/jtacoma/gozmqutil"
    )

[ØMQ](http://www.zeromq.org) abstractions for [Go](http://golang.org).

This package aims to implement, first, the [recommended binding abstractions](http://www.zeromq.org/topics:binding-abstractions).

## Status

This project is in early development, the public API may change significantly before 1.0.  It is not yet recommended for use production (though [ØMQ](http://www.zeromq.org) and [gozmq](https://github.com/alecthomas/gozmq) are).

## Building

Build tags are used to distinguish versions of ØMQ.  Version 2.1 is `zmq_2_1`, 2.2 is `zmq_2_x`, and 3.2 is `zmq_3_x`.  The package will not build unless you specify one of these tags.

## Implemented Features

### Smarter Context

A context remembers its sockets and, when the context is closed, it will close all its sockets too.  This includes a context-scoped linger option, applied to all sockets just before the context closes them.

### Friendlier Sockets

All options supported in [gozmq](https://github.com/alecthomas/gozmq) are available here through option-specific getter/setter methods.

### Reactor/Device Loop

A reactor loop is provided for socket events.

## Planned Features

* Bind method that allows "bind to random free port" (and returns resulting port).

## License

Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
