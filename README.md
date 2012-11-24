gzmq
====

This package lets messages be received from and sent to [ØMQ](http://www.zeromq.org) sockets through [Go](http://golang.org) channels.

The gzmq package is based on [gozmq](https://github.com/alecthomas/gozmq).

For examples and API documentation, see [go.pkgdoc.org](http://go.pkgdoc.org/github.com/jtacoma/gzmq).

Purpose
-------

ØMQ and Go have a few things in common: both aim to be minimal in order to do a few things very well, and both have been designed to support concurrency through communication.

*A Go application that uses ØMQ sockets should be a joy to develop.*  Unfortunately, ØMQ doesn't quite fit in with the way things are done in Go.

ØMQ sockets are not thread-safe.  This is perfectly acceptable when ØMQ enables message passing between threads for languages that don't already provide such an abstraction, but Go does.  Go's channels are thread-safe.  They're also syntactically sweet and easy to use.

This package implements channel-based sending and receiving for ØMQ sockets.

Problems
--------

Buffered channels hold message outside the scope of ØMQ's own buffers.  This means that if you queue some outgoing messages in a channel, then close the socket before they've all been sent, ØMQ will (by default) not close the socket until all the messages it knows about have been sent.  In this case, messages waiting in the channel will not be sent.  The current version of gzmq does not address this issue.

Channels do not provide all the options available when sending/receiving through ØMQ's API.  For example, using ØMQ directly you can do a non-blocking send that fails if the message can't immediately be sent.  This can still be accomplished from code that has access to the Polling (through its Sync method) but doing so lessens the benefit of using gzmq.

License
-------

The *gzmq* package is open source under the [MIT License](http://opensource.org/licenses/MIT).
