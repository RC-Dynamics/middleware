package main

import (
	"net"
)

type ClientHandlerTCP struct {
	address string
	conn    net.Conn
}

func (handler *ClientHandlerTCP) connect(address string) {
	conn, err := net.Dial("tcp", address)
	checkError(err)
	handler.address = address
	handler.conn = conn
}

func (handler *ClientHandlerTCP) read(size int) []byte {
	buffer := make([]byte, size)
	_, err := handler.conn.Read(buffer)
	checkError(err)
	return buffer
}

func (handler *ClientHandlerTCP) send(buffer []byte) {
	_, err := handler.conn.Write(buffer)
	checkError(err)
}

func (handler *ClientHandlerTCP) close() {
	handler.conn.Close()
}
