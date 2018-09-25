package main

import "net"

type ServerHandler struct {
	tp      string
	port    int
	conn    net.Conn
	addr    net.Addr
	handler interface{}
}

func (server *ServerHandler) create() {
	switch server.tp {
	case "tcp":
		server.handler = ServerHandlerTCP{server.port}
		// conn = handlerTCP.create(port)
	case "udp":
		server.handler = ServerHandlerUDP{server.port}
		// conn = handlerUDP.crete(port)
	}
	server.conn = server.handler.create()
}

func (server *ServerHandler) read(int size) {
	buff, addr := server.handler.read(size, conn)
	server.addr = addr
	return buff
}

func (server *ServerHandler) send(buff []byte) {
	server.handler.send(buff, server.conn, server.addr)
}

func (server *ServerHandler) close() {
	server.conn.Close()
}
