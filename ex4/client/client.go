package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	address := "localhost:8080"

	clientRequestHandler := ClientRequestHandler{"rpc", nil}
	clientRequestHandler.connect(address)

	for _, input := range []string{"rpc-1k-middleware.csv"} {
		for _, qtd := range []int{1000} {
			filename := input
			fmt.Println(filename)
			file, err := os.Create(filename)
			checkError(err)
			defer file.Close()
			// BechMarket here
			for i := 0; i < qtd; i++ {
				time1 := time.Now()

				// Application
				clientRequestHandler.send([]byte("HELLOWORLD"))
				clientRequestHandler.read(10)

				time2 := time.Now()
				// fmt.Fprintf(os.Stderr, "%s\n", string(clientRequestHandler.read(10)))
				elapsedTime := float64(time2.Sub(time1).Nanoseconds()) / 1000000
				fmt.Fprintln(file, elapsedTime)
				checkError(err)
				time.Sleep(10 * time.Millisecond)
				// To here

			}

		}
	}
	clientRequestHandler.close()

	os.Exit(0)
}
