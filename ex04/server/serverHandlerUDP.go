package main

import (
	"fmt"
	"net"
	"os"
)

type ServerHandlerUDP struct {
	port int
}

func (handler *ServerHandlerUDP) create() {
	conn, err := net.ListenPacket("udp", ":"+string(handler.port))
	handler.checkError(err)
	return conn
}

func (handler *ServerHandlerUDP) read(size int, conn net.Conn) []byte {
	buffer := make([]byte, size)
	_, addr, err := conn.ReadFrom(buffer)
	handler.checkError(err)
	return buffer, addr
}

func (handler *ServerHandlerUDP) send(buffer []byte, addr net.Addr, conn net.Conn) {
	_, err = conn.WriteTo(buffer, addr)
	handler.checkError(err)
}

func (handler *ServerHandlerUDP) checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
