package gzmq

import (
	"errors"

	zmq "github.com/alecthomas/gozmq"
)

type Bridge struct {
	context zmq.Context
}

func NewBridge(context zmq.Context) (*Bridge, error) {
	var err error
	if context == nil {
		context, err = zmq.NewContext()
	}
	if err != nil {
		return nil, err
	}
	return &Bridge{context: context}, nil
}

func (b *Bridge) Context() zmq.Context {
	return b.context
}

func (b *Bridge) RegisterIn(s zmq.Socket) (<-chan [][]byte, error) {
	return nil, errors.New("TODO: implement this.")
}

func (b *Bridge) RegisterOut(s zmq.Socket) (chan<- [][]byte, error) {
	return nil, errors.New("TODO: implement this.")
}

func (b *Bridge) RegisterInOut(s zmq.Socket) (<-chan [][]byte, chan<- [][]byte, error) {
	return nil, nil, errors.New("TODO: implement this.")
}

func (b *Bridge) Start() {
	panic("TODO: implement this.")
}

func (b *Bridge) Stop() {
	panic("TODO: implement this.")
}
