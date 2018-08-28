package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

// BUFFERSIZE for file transfer
const BUFFERSIZE = 1024

// Main Server
func main() {
	service := ":8080"

	listener, err := net.Listen("tcp", service)
	// listener, err := net.Listen("udp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		go handleClientTCP(conn)
	}
}

func handleClientTCP(conn net.Conn) {
	defer conn.Close()
	// Getting File Name
	bufferFileName := make([]byte, 50)
	_, err := conn.Read(bufferFileName)
	checkError(err)
	fileName := strings.Trim(string(bufferFileName), ":")
	fmt.Println("FileName: ", fileName)

	// Read File and Get Its Size!
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	fileInfo, err := file.Stat()
	checkError(err)

	// Sending Size
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 15)
	_, err = conn.Write([]byte(fileSize))
	checkError(err)

	// Sending File:
	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}
		conn.Write(sendBuffer)
	}
	// we're finished with this client
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}
