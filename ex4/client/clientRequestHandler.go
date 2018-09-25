package main

import (
	"fmt"
	"os"
)

type Handler interface {
	connect(string)
	read(int) []byte
	send([]byte)
	close()
}

type ClientRequestHandler struct {
	tp      string
	handler Handler
}

func (client *ClientRequestHandler) connect(address string) {
	switch client.tp {
	case "tcp":
		client.handler = &ClientHandlerTCP{"", nil}
	case "udp":
		client.handler = &ClientHandlerUDP{"", nil}
	}
	client.handler.connect(address)
}

func (client *ClientRequestHandler) read(size int) []byte {
	return client.handler.read(size)
}

func (client *ClientRequestHandler) send(buff []byte) {
	client.handler.send(buff)
}

func (client *ClientRequestHandler) close() {
	client.handler.close()
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
