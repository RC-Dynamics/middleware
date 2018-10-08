package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	address := "localhost:8080"

	client_proxy := ClientProxy{address}
	string_proxy := StringProxy{client_proxy}

	for _, qtd := range []int{1000, 5000, 10000} {
		filename := "tcp-" + strconv.Itoa(qtd/1000) + "k-middleware.csv"
		fmt.Println(filename)
		file, err := os.Create(filename)
		checkError(err)
		defer file.Close()
		// BechMarket here
		for i := 0; i < qtd; i++ {
			time1 := time.Now()

			// Application
			string_proxy.toUpper("helloworld")
			// fmt.Println(string_proxy.toUpper("helloworld"))

			time2 := time.Now()

			elapsedTime := float64(time2.Sub(time1).Nanoseconds()) / 1000000
			fmt.Fprintln(file, elapsedTime)
			checkError(err)
			time.Sleep(10 * time.Millisecond)
			// To here

		}

	}

	os.Exit(0)
}
