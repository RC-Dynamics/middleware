package main

import (
	"strings"
)

// Main Server
func main() {
	serverRequestHandler := ServerHandler{"tcp", ":8080", nil}
	serverRequestHandler.create()
	for {
		data := string(serverRequestHandler.read(10))
		serverRequestHandler.send([]byte(strings.ToLower(data)))
		// fmt.Fprintf(os.Stderr, "%s\n", data)
	}
	serverRequestHandler.close()
}
