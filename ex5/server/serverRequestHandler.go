package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
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
		server.handler = &ServerHandlerTCP{server.port, nil, nil}
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
	return strings.ToUpper(str)
}

func (proxy *StringProxy) toLower(str string) string {
	return strings.ToLower(str)
}

func invoke(clientProxy ClientProxy) Termination {
	srh := ServerHandler{
		tp:   "tcp",
		port: clientProxy.Address,
	}

	stringProxy := StringProxy{}
	ter := Termination{}

	for {
		srh.create()

		msgToBeUnmarshalled := srh.read(500)
		msgToBeUnmarshalled = bytes.Trim(msgToBeUnmarshalled, "\x00")

		var msgUnmarshalled Message

		json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)
		// fmt.Print(msgToBeUnmarshalled)
		// fmt.Println(" - " + msgUnmarshalled.Body.RequestHeader.Operation)

		if msgUnmarshalled.Body.RequestHeader.Operation == "toUpper" {
			ter.Result = stringProxy.toUpper(msgUnmarshalled.Body.RequestBody.Parameters[0])
		} else if msgUnmarshalled.Body.RequestHeader.Operation == "toLower" {
			ter.Result = stringProxy.toLower(msgUnmarshalled.Body.RequestBody.Parameters[0])
		}

		replyHeader := ReplyHeader{
			ServiceContext: "",
			Id:             0,
			Status:         0,
		}

		replyBody := ReplyBody{
			Result: ter.Result,
		}

		messageHeader := MessageHeader{
			Magic:       "protocolo",
			Version:     0,
			ByteOrders:  false,
			MessageType: 0,
			MessageSize: 0,
		}

		messageBody := MessageBody{
			ReplyHeader: replyHeader,
			ReplyBody:   replyBody,
		}

		msgtoBeMarshalled := Message{
			Header: messageHeader,
			Body:   messageBody,
		}

		msgMarshalled, err := json.Marshal(msgtoBeMarshalled)
		checkError(err)

		time.Sleep(10 * time.Millisecond)

		srh.send(msgMarshalled)

		srh.close()

	}
}
