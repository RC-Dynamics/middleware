package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "../protobuff"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	name    = "hello world"
)

func main() {
	log.Print("Starting Client Test")
	for _, qtd := range []int{1000, 5000, 10000} {
		filename := "grpc-" + strconv.Itoa(qtd/1000) + "k.csv"
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
			c := pb.NewStringManipulationClient(conn)

			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)

			_, err = c.Upper(ctx, &pb.StrRequest{Name: name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}

			cancel()
			conn.Close()
			time2 := time.Now()
			elapsedTime := float64(time2.Sub(time1).Nanoseconds()) / 1000000
			fmt.Fprintln(file, elapsedTime)
			time.Sleep(10 * time.Millisecond)
		}
		log.Printf("End Test of: %d", qtd)
	}

}
