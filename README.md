# gzmq

[ØMQ](http://www.zeromq.org) abstractions for [Go](http://golang.org).

This package aims to implement, first, the [recommended binding abstractions](http://www.zeromq.org/topics:binding-abstractions).

## Features

* A context remembers its sockets and, when the context is closed, it will close all its sockets too.
* `LINGER` is context option, applied to its sockets.
