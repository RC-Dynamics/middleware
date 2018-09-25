package main

import (
	"fmt"
	"os"
	"strings"
)

// Main Server
func main() {
	serverRequestHandler := ServerHandler{"tcp", ":8080", nil}
	serverRequestHandler.create()
	data := string(serverRequestHandler.read(10))
	fmt.Fprintf(os.Stderr, "%s\n", data)
	serverRequestHandler.send([]byte(strings.ToLower(data)))
	serverRequestHandler.close()
}
