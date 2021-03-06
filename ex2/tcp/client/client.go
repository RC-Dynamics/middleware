package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// BUFFERSIZE for file transfer
const BUFFERSIZE = 1024

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]

	for _, input := range []string{"test-1KB.txt", "test-1MB.txt"} {
		for _, qtd := range []int{1000, 5000, 10000} {
			filename := "result" + input[4:8] + "-" + strconv.Itoa(qtd/1000) + "k.csv"
			fmt.Println(filename)
			file, err := os.Create(filename)
			checkError(err)
			defer file.Close()
			// BechMarket here
			for i := 0; i < qtd; i++ {
				time1 := time.Now()
				conn, err := net.Dial("tcp", address)
				checkError(err)
				requestFileTCP(input, conn)
				conn.Close()
				time2 := time.Now()
				elapsedTime := float64(time2.Sub(time1).Nanoseconds()) / 1000000
				fmt.Fprintln(file, elapsedTime)
				checkError(err)
				time.Sleep(10 * time.Millisecond)
				// To here
			}

		}
	}

	os.Exit(0)
}

func requestFileTCP(fileName string, conn net.Conn) {
	// Sending File Name
	_, err := conn.Write([]byte(fillString(fileName, 50)))
	checkError(err)

	// Getting File Size
	bufferFileSize := make([]byte, 15)
	_, err = conn.Read(bufferFileSize)
	checkError(err)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	// fmt.Println("FileSize: ", fileSize)

	// Getting File:
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()

	var recSize int64
	recSize = 0
	for {
		if (fileSize - recSize) < BUFFERSIZE {
			io.CopyN(file, conn, (fileSize - recSize))
			conn.Read(make([]byte, (recSize+BUFFERSIZE)-fileSize))
			recSize = fileSize
			break
		}
		io.CopyN(file, conn, BUFFERSIZE)
		recSize += BUFFERSIZE
		if recSize == fileSize {
			break
		}
	}

	// fmt.Println("File Received: ", fileName, "  Size: ", recSize)

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
