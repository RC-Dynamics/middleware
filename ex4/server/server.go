package main

import (
	"fmt"
	"os"
)

// Main Server
func main() {
	serverRequestHandler := ServerHandler{"tcp", 8080, nil, nil, nil}
	serverRequestHandler.create()
	fmt.Fprintf(os.Stderr, "%s", string(serverRequestHandler.read(10)))
	serverRequestHandler.close()
}
