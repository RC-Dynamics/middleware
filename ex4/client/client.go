package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	address := "localhost:8080"

	clientRequestHandler := ClientRequestHandler{"tcp", nil}
	clientRequestHandler.connect(address)

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

				// Application
				clientRequestHandler.send([]byte("HELLOWORLD"))

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
