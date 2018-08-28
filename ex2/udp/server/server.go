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
	for {
		pc, err := net.ListenPacket("udp", service)
		checkError(err)
		handleClientUDP(pc)
	}

	// for {
	// 	conn, err := listener.Accept()

	// 	if err != nil {
	// 		continue
	// 	}
	// 	go handleClientUDP(conn)
	// }
}

func handleClientUDP(conn net.PacketConn) {
	defer conn.Close()
	// Getting File Name
	bufferFileName := make([]byte, 50)
	_, addr, err := conn.ReadFrom(bufferFileName)
	checkError(err)
	fileName := strings.Trim(string(bufferFileName), ":")

	// Read File and Get Its Size!
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	fileInfo, err := file.Stat()
	checkError(err)

	// Sending Size
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 15)
	_, err = conn.WriteTo([]byte(fileSize), addr)
	checkError(err)

	fmt.Println("FileName: ", fileName, "   Size: ", fileInfo.Size())

	// Sending File:
	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}
		conn.WriteTo(sendBuffer, addr)
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
