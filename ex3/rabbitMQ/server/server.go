package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/streadway/amqp"
)

// BUFFERSIZE for file transfer
const BUFFERSIZE = 1024

// Main Server
func main() {

	// service := ":8080"
	// listener, err := net.Listen("tcp", service)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	checkError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	checkError(err)
	defer ch.Close()

	fileQueue, err := ch.QueueDeclare(
		"file", // name
		false,  // durable
		false,  // delete when usused
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

	requestCh, err := ch.Consume(
		requestQueue.Name, // queue
		"",                // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)

	for {
		handleClient(ch, requestCh, fileQueue)
	}
}

func publish(ch *amqp.Channel, queue amqp.Queue, output []byte) {
	ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        (output),
		})
}

func handleClient(ch *amqp.Channel, requestCh <-chan amqp.Delivery, fileQ amqp.Queue) {
	// Getting File Name
	// bufferFileName := make([]byte, 50)
	// _, err := conn.Read(bufferFileName)
	// checkError(err)
	// fileName := strings.Trim(string(bufferFileName), ":")

	fileName := string((<-requestCh).Body)
	// fmt.Println("FileName:  ", fileName)

	// // Read File and Get Its Size!
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()
	fileInfo, err := file.Stat()
	checkError(err)

	// // Sending Size
	// fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 15)
	// _, err = conn.Write([]byte(fileSize))
	// checkError(err)

	publish(ch, fileQ, []byte(string(strconv.FormatInt(fileInfo.Size(), 10))))
	fmt.Println("Sending -> Name: ", fileName, "   Size: ", string(strconv.FormatInt(fileInfo.Size(), 10)))

	// // Sending File:
	sendBuffer := make([]byte, BUFFERSIZE)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		} else {
			checkError(err)
		}
		publish(ch, fileQ, sendBuffer)
		// conn.Write(sendBuffer)
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
