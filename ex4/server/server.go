package main

import (
	"fmt"
	"os"
)

// Main Server
func main() {
	serverRequestHandler := ServerHandler{"tcp", ":8080", nil}
	serverRequestHandler.create()
	fmt.Fprintf(os.Stderr, "%s\n", string(serverRequestHandler.read(10)))
	serverRequestHandler.send([]byte("VEM ni MIM"))
	serverRequestHandler.close()
}
