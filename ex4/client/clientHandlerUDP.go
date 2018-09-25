package main

import (
	"net"
)

type ClientHandlerUDP struct {
	address string
	conn    net.Conn
}

func (handler *ClientHandlerUDP) connect(address string) {
	conn, err := net.Dial("udp", address)
	checkError(err)
	handler.address = address
	handler.conn = conn
}

func (handler *ClientHandlerUDP) read(size int) []byte {
	buffer := make([]byte, size)
	_, err := handler.conn.Read(buffer)
	checkError(err)
	return buffer
}

func (handler *ClientHandlerUDP) send(buffer []byte) {
	_, err := handler.conn.Write(buffer)
	checkError(err)
}

func (handler *ClientHandlerUDP) close() {
	handler.conn.Close()
}
