package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	address := "localhost:8080"

	for _, mid_type := range []string{"tcp"} {

		for _, qtd := range []int{1000, 5000, 10000} {
			filename := mid_type + "-" + strconv.Itoa(qtd/1000) + "k-middleware.csv"
			fmt.Println(filename)
			file, err := os.Create(filename)
			checkError(err)
			defer file.Close()
			// BechMarket here
			for i := 0; i < qtd; i++ {
				time1 := time.Now()
				clientRequestHandler := ClientRequestHandler{mid_type, nil}
				clientRequestHandler.connect(address)

				// Application
				clientRequestHandler.send([]byte("helloworld"))
				clientRequestHandler.read(500)

				clientRequestHandler.close()
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

	os.Exit(0)
}
