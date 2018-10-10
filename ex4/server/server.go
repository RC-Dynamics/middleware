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
		data := serverRequestHandler.read(10)
		// data = bytes.Trim(data, "\x00")
		sData := string(data)
		time.Sleep(10 * time.Millisecond)
		serverRequestHandler.send([]byte(strings.ToUpper(sData)))
		// fmt.Fprintf(os.Stderr, "%s\n", data)
		serverRequestHandler.close()
	}
}
