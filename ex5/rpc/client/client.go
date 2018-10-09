package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	for _, qtd := range []int{1} { //, 5000, 10000} {
		filename := "rpc-" + strconv.Itoa(qtd/1000) + "k.csv"
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		fmt.Println(filename)

		for i := 0; i < qtd; i++ {
			time1 := time.Now()
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				log.Fatal("dialing:", err)
			}
			c := pb.NewGreeterClient(conn)

			// Contact the server and print out its response.
			name := defaultName
			if len(os.Args) > 1 {
				name = os.Args[1]
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)

			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.Message)

			cancel()
			conn.Close()
			time2 := time.Now()
			elapsedTime := float64(time2.Sub(time1).Nanoseconds()) / 1000000
			fmt.Fprintln(file, elapsedTime)
			time.Sleep(10 * time.Millisecond)
		}
	}

}
