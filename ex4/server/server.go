package main

import (
	"strings"
	"time"
)

// Main Server
func main() {
	for {
		serverRequestHandler := ServerHandler{"tcp", ":8080", nil}
		serverRequestHandler.create()
		data := string(serverRequestHandler.read(10))
		time.Sleep(10 * time.Millisecond)
		serverRequestHandler.send([]byte(strings.ToLower(data)))
		// fmt.Fprintf(os.Stderr, "%s\n", data)
		serverRequestHandler.close()
	}
}
