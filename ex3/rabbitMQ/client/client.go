package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
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
				conn, err := amqp.Dial("amqp://guest:guest@" + address)
				checkError(err)
				ch, err := conn.Channel()
				checkError(err)

				fileQueue, _ := ch.QueueDeclare(
					"file", // name
					false,  // durable
					false,  // delete when unused
					false,  // exclusive
					false,  // no-wait
					nil,    // arguments
				)

				requestQueue, err := ch.QueueDeclare(
					"request", // name
					false,     // durable
					false,     // delete when usused
					false,     // exclusive
					false,     // no-wait
					nil,       // arguments
				)

				requestFile(input, ch, fileQueue, requestQueue)

				time2 := time.Now()
				ch.Close()
				conn.Close()
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

func requestFile(fileName string, ch *amqp.Channel, fileQ amqp.Queue, requestQ amqp.Queue) {
	// Getting File Size
	fileCh, err := ch.Consume(
		fileQ.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	checkError(err)

	// Sending File Name
	err = ch.Publish(
		"",            // exchange
		requestQ.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fileName),
		})
	// _, err := conn.Write([]byte(fillString(fileName, 50)))
	checkError(err)

	fileSize, _ := (strconv.ParseInt(string((<-fileCh).Body), 10, 64))
	// fmt.Println("FileSize: ", (fileSize))

	// bufferFileSize := make([]byte, 15)
	// _, err = conn.Read(bufferFileSize)
	// checkError(err)
	// fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	// fmt.Println("FileSize: ", fileSize)

	// Getting File:
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()

	var recSize int64
	recSize = 0
	for {
		if (fileSize - recSize) < BUFFERSIZE {
			// recBuffer[:(fileSize - recSize)]
			file.Write((<-fileCh).Body[:(fileSize - recSize)])
			recSize = fileSize
			break
		}
		file.Write((<-fileCh).Body)
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
