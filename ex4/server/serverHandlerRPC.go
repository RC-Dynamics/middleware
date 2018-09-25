package main

import (
	"net"
	"net/http"
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

type ServerHandlerRPC struct {
	port string
}

func (handler *ServerHandlerRPC) create() {
	rpc.Register(new(Str))
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", handler.port)
	checkError(e)
	http.Serve(l, nil)
}

func (handler *ServerHandlerRPC) read(size int) []byte {
	return nil
}

func (handler *ServerHandlerRPC) send(buffer []byte) {

}

func (handler *ServerHandlerRPC) close() {
}
