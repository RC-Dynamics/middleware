package main

import (
	"fmt"
	"net"
	"os"
)

type ServerHandlerTCP struct {
	port string
	conn net.Conn
}

func (handler ServerHandlerTCP) create() {
	listener, err := net.Listen("tcp", handler.port)
	checkError(err)
	conn, err := listener.Accept()
	handler.conn = conn
}

func (handler ServerHandlerTCP) read(size int) []byte {
	buffer := make([]byte, size)
	_, err := handler.conn.Read(buffer)
	checkError(err)
	return buffer
}

func (handler ServerHandlerTCP) send(buffer []byte) {
	_, err := handler.conn.Write(buffer)
	checkError(err)
}

func (handler ServerHandlerTCP) close() {
	handler.conn.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
