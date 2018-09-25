package main

import (
	"fmt"
	"os"
)

// Main Server
func main() {
	serverRequestHandler := ServerHandler{"udp", ":8080", nil}
	serverRequestHandler.create()
	fmt.Fprintf(os.Stderr, "%s", string(serverRequestHandler.read(10)))
	serverRequestHandler.close()
}
