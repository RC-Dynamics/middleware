package main

import (
	"fmt"
	"net"
	"os"
)

type ServerHandlerTCP struct {
	port int
}

func (handler *ServerHandlerTCP) create() {
	listener, err := net.Listen("tcp", ":"+string(handler.port))
	handler.checkError(err)
	conn, err := listener.Accept()
	return conn
}

func (handler *ServerHandlerTCP) read(size int, conn net.Conn) []byte {
	buffer := make([]byte, size)
	_, err := conn.Read(buffer)
	handler.checkError(err)
	return buffer, _
}

func (handler *ServerHandlerTCP) send(buffer []byte, conn net.Conn) {
	_, err = conn.Write(buffer)
	handler.checkError(err)
}

func (handler *ServerHandlerTCP) checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
