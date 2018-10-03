package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"time"

	"./server"
)

func main() {
	for _, qtd := range []int{1000, 5000, 10000} {
		filename := "rpc-" + strconv.Itoa(qtd/1000) + "k.csv"
		file, err := os.Create(filename)
		checkError(err)
		fmt.Println(filename)

		client, err := rpc.DialHTTP("tcp", "localhost:8081")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		for i := 0; i < qtd; i++ {
			time1 := time.Now()

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

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
