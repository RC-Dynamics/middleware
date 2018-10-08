package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
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
		// case "udp":
		// 	client.handler = &ClientHandlerUDP{"", nil}
		// case "rpc":
		// 	client.handler = &ClientHandlerRPC{"", nil, nil}
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
	inv := Invocation{
		address:     proxy.client_prxy.address,
		method_name: "toUpper",
		parameters:  []string{str},
	}
	ter := invoke(inv)
	return ter.result
}

func (proxy *StringProxy) toLower(str string) string {
	inv := Invocation{
		address:     proxy.client_prxy.address,
		method_name: "toLower",
		parameters:  []string{str},
	}
	ter := invoke(inv)
	return ter.result
}

func invoke(inv Invocation) Termination {
	chr := ClientRequestHandler{
		tp: "tcp",
	}
	chr.connect(inv.address)
	reqHeader := RequestHeader{
		context:          "",
		id:               0,
		responseExpected: true,
		objectKey:        0,
		operation:        inv.method_name,
	}

	reqBody := RequestBody{
		parameters: inv.parameters,
	}

	messageHeader := MessageHeader{
		magic:       "MIOP",
		version:     0,
		byteOrders:  false,
		messageType: 0,
		messageSize: 0,
	}

	messageBody := MessageBody{
		requestHeader: reqHeader,
		requestBody:   reqBody,
	}

	msgtoBeMarshalled := Message{
		header: messageHeader,
		body:   messageBody,
	}

	msgMarshalled, _ := json.Marshal(msgtoBeMarshalled)

	chr.send(msgMarshalled)

	msgToBeUnmarshalled := chr.read(10)

	var msgUnmarshalled Message

	json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)

	ter := Termination{
		result: msgUnmarshalled.body.replybody.result,
	}

	return ter
}
