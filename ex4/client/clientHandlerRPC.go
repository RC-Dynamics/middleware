package main

import (
	"net/rpc"
	"strings"
)

type Args struct {
	Input []byte
}

type Str string

func (t *Str) Lower(args *Args, reply *[]byte) error {
	*reply = []byte(strings.ToLower(string(args.Input)))
	return nil
}

type ClientHandlerRPC struct {
	address string
	client  *rpc.Client
	reply   []byte
}

func (handler *ClientHandlerRPC) connect(address string) {
	client, err := rpc.DialHTTP("tcp", address)
	checkError(err)
	handler.address = address
	handler.client = client
}

func (handler *ClientHandlerRPC) read(size int) []byte {
	return handler.reply
}

func (handler *ClientHandlerRPC) send(buffer []byte) {
	args := &Args{buffer}
	var reply []byte
	err := handler.client.Call("Str.Lower", args, &reply)
	handler.reply = reply
	checkError(err)
}

func (handler *ClientHandlerRPC) close() {
}
