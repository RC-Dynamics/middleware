package main

import (
	"fmt"
	"net"
)

type ServerHandlerUDP struct {
	port string
	addr net.Addr
	conn net.PacketConn
}

func (handler ServerHandlerUDP) create() {
	conn, err := net.ListenPacket("udp", handler.port)
	checkError(err)
	handler.conn = conn
	fmt.Println(handler.conn)
}

func (handler ServerHandlerUDP) read(size int) []byte {
	buffer := make([]byte, size)
	_, addr, err := handler.conn.ReadFrom(buffer)
	handler.addr = addr
	checkError(err)
	return buffer
}

func (handler ServerHandlerUDP) send(buffer []byte) {
	_, err := handler.conn.WriteTo(buffer, handler.addr)
	checkError(err)
}

func (handler ServerHandlerUDP) close() {
	handler.conn.Close()
}
