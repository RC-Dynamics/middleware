package main

import (
	"bytes"
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
	Address string
}

type StringProxy struct {
	Client_proxy ClientProxy
}

type Invocation struct {
	Address     string
	Method_name string
	Parameters  []string
}

type Termination struct {
	Result string
}

type MessageHeader struct {
	Magic       string
	Version     int
	ByteOrders  bool
	MessageType int
	MessageSize int
}

type RequestHeader struct {
	Context          string
	Id               int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type RequestBody struct {
	Parameters []string
}

type ReplyHeader struct {
	ServiceContext string
	Id             int
	Status         int
}

type ReplyBody struct {
	Result string
}

type MessageBody struct {
	RequestHeader RequestHeader
	RequestBody   RequestBody
	ReplyHeader   ReplyHeader
	ReplyBody     ReplyBody
}

type Message struct {
	Header MessageHeader
	Body   MessageBody
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (proxy *StringProxy) toUpper(str string) string {
	inv := Invocation{
		Address:     proxy.Client_proxy.Address,
		Method_name: "toUpper",
		Parameters:  []string{str},
	}
	ter := invoke(inv)
	return ter.Result
}

func (proxy *StringProxy) toLower(str string) string {
	inv := Invocation{
		Address:     proxy.Client_proxy.Address,
		Method_name: "toLower",
		Parameters:  []string{str},
	}
	ter := invoke(inv)
	return ter.Result
}

func invoke(inv Invocation) Termination {
	chr := ClientRequestHandler{
		tp: "tcp",
	}

	chr.connect(inv.Address)

	reqHeader := RequestHeader{
		Context:          "",
		Id:               0,
		ResponseExpected: true,
		ObjectKey:        0,
		Operation:        inv.Method_name,
	}

	reqBody := RequestBody{
		Parameters: inv.Parameters,
	}

	messageHeader := MessageHeader{
		Magic:       "MIOP",
		Version:     0,
		ByteOrders:  false,
		MessageType: 0,
		MessageSize: 0,
	}

	messageBody := MessageBody{
		RequestHeader: reqHeader,
		RequestBody:   reqBody,
	}

	msgtoBeMarshalled := Message{
		Header: messageHeader,
		Body:   messageBody,
	}

	msgMarshalled, err := json.Marshal(msgtoBeMarshalled)
	checkError(err)

	chr.send(msgMarshalled)

	msgToBeUnmarshalled := chr.read(500)
	msgToBeUnmarshalled = bytes.Trim(msgToBeUnmarshalled, "\x00")

	var msgUnmarshalled Message

	json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)

	ter := Termination{
		Result: msgUnmarshalled.Body.ReplyBody.Result,
	}

	chr.close()

	return ter
}
