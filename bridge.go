package gzmq

import (
	zmq "github.com/alecthomas/gozmq"
)

type Bridge struct {
	Context zmq.Context
}

func (b *Bridge) RegisterIn(s zmq.Socket) (<-chan [][]byte, error) {
	panic("TODO: implement this.")
}

func (b *Bridge) RegisterOut(s zmq.Socket) (chan<- [][]byte, error) {
	panic("TODO: implement this.")
}

func (b *Bridge) RegisterInOut(s zmq.Socket) (<-chan [][]byte, chan<- [][]byte) {
	panic("TODO: implement this.")
}

func (b *Bridge) Start() {
	panic("TODO: implement this.")
}

func (b *Bridge) Stop() {
	panic("TODO: implement this.")
}
