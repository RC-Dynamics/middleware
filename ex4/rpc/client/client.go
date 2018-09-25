package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"

	"./server"
)

func main() {
	filename := "rcp-1k.txt"
	file, err := os.Create(filename)
	checkError(err)
	for i := 0; i < 1000; i++ {
		time1 := time.Now()

		client, err := rpc.DialHTTP("tcp", "localhost:8081")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		// Synchronous call
		args := &server.Args{"HELLOWORLD"}
		var reply string
		err = client.Call("Str.Lower", args, &reply)
		if err != nil {
			log.Fatal("call error:", err)
		}
		time2 := time.Now()
		elapsedTime := float64(time2.Sub(time1).Nanoseconds()) / 1000000
		fmt.Fprintln(file, elapsedTime)

		time.Sleep(10 * time.Millisecond)
		// fmt.Printf("Lower: %s = %s\n", args.Txt, reply)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
