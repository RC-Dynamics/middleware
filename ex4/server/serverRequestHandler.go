package main

type Handler interface {
	create()
	read(int) []byte
	send([]byte)
	close()
}

type ServerHandler struct {
	tp      string
	port    string
	handler Handler
}

func (server *ServerHandler) create() {
	switch server.tp {
	case "tcp":
		server.handler = &ServerHandlerTCP{server.port, nil}
		// conn = handlerTCP.create(port)
	case "udp":
		server.handler = &ServerHandlerUDP{server.port, nil, nil}
		// conn = handlerUDP.crete(port)
	}
	server.handler.create()
}

func (server *ServerHandler) read(size int) []byte {
	return server.handler.read(size)
}

func (server *ServerHandler) send(buff []byte) {
	server.handler.send(buff)
}

func (server *ServerHandler) close() {
	server.handler.close()
}
