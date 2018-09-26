package main

import (
	"strings"
)

// Main Server
func main() {
	serverRequestHandler := ServerHandler{"rpc", ":8080", nil}
	serverRequestHandler.create()
	for {
		data := string(serverRequestHandler.read(10))
		// fmt.Fprintf(os.Stderr, "%s\n", data)
		serverRequestHandler.send([]byte(strings.ToLower(data)))
	}
	serverRequestHandler.close()
}
