package main

import (
	"fmt"
	"log"
	"net/rpc"

	"./server"
)

func main() {
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
	fmt.Printf("Lower: %s = %s\n", args.Txt, reply)
}
