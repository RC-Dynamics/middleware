package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// BUFFERSIZE for file transfer
const BUFFERSIZE = 1024

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]

	conn, err := net.Dial("udp", address)
	checkError(err)
	defer conn.Close()

	requestFileUDP("123.JPG", conn)

	os.Exit(0)
}

func requestFileUDP(fileName string, conn net.Conn) {
	// Sending File Name
	_, err := conn.Write([]byte(fillString(fileName, 50)))
	checkError(err)

	// Getting File Size
	bufferFileSize := make([]byte, 15)
	_, err = conn.Read(bufferFileSize)
	checkError(err)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	fmt.Println("FileSize: ", fileSize)

	// Getting File:
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()

	var recSize int64
	recSize = 0
	recBuffer := make([]byte, BUFFERSIZE)
	for {
		conn.Read(recBuffer)
		if (fileSize - recSize) < BUFFERSIZE {
			file.WriteAt(recBuffer, (fileSize - recSize))
			recSize = fileSize
			break
		}
		fmt.Println(recSize)
		file.WriteAt(recBuffer, BUFFERSIZE)
		recSize += BUFFERSIZE
		if recSize == fileSize {
			break
		}
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
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
