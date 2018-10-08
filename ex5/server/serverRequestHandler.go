package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

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
		// case "udp":
		// 	server.handler = &ServerHandlerUDP{server.port, nil, nil}
		// 	// conn = handlerUDP.crete(port)
		// case "rpc":
		// 	server.handler = &ServerHandlerRPC{server.port}
		// 	// conn = handlerUDP.crete(port)
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

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}

type ServerFunc interface {
	toUpper(string) string
	toLower(string) string
}

type ClientProxy struct {
	address string
}

type StringProxy struct {
	client_prxy ClientProxy
}

type Invocation struct {
	address     string
	method_name string
	parameters  []string
}

type Termination struct {
	result string
}

type MessageHeader struct {
	magic       string
	version     int
	byteOrders  bool
	messageType int
	messageSize int
}

type RequestHeader struct {
	context          string
	id               int
	responseExpected bool
	objectKey        int
	operation        string
}

type RequestBody struct {
	parameters []string
}

type ReplyHeader struct {
	serviceContext string
	id             int
	status         int
}

type ReplyBody struct {
	result string
}

type MessageBody struct {
	requestHeader RequestHeader
	requestBody   RequestBody
	replyHeader   ReplyHeader
	replyBody     ReplyBody
}

type Message struct {
	header MessageHeader
	body   MessageBody
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (proxy *StringProxy) toUpper(str string) string {
	return strings.ToUpper(str)
}

func (proxy *StringProxy) toLower(str string) string {
	return strings.ToLower(str)
}

func invoke(clientProxy ClientProxy) Termination {
	srh := ServerHandler{
		tp:   "tcp",
		port: clientProxy.address,
	}
	srh.create()

	stringProxy := StringProxy{}
	ter := Termination{}

	for {
		msgToBeUnmarshalled := srh.read(10)

		var msgUnmarshalled Message

		json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)

		// TODO: processar
		if msgUnmarshalled.body.requestHeader.operation == "toUpper" {
			stringProxy.toUpper(msgUnmarshalled.body.requestBody.parameters[0])
		} else if msgUnmarshalled.body.requestHeader.operation == "toLower" {
			stringProxy.toLower(msgUnmarshalled.body.requestBody.parameters[0])
		}

		replyHeader := ReplyHeader{
			serviceContext: "",
			id:             0,
			status:         0,
		}

		replyBody := ReplyBody{
			result: ter.result,
		}

		messageHeader := MessageHeader{
			magic:       "protocolo",
			version:     0,
			byteOrders:  false,
			messageType: 0,
			messageSize: 0,
		}

		messageBody := MessageBody{
			replyHeader: replyHeader,
			replyBody:   replyBody,
		}

		msgtoBeMarshalled := Message{
			header: messageHeader,
			body:   messageBody,
		}

		msgMarshalled, _ := json.Marshal(msgtoBeMarshalled)

		srh.send(msgMarshalled)

	}
}
