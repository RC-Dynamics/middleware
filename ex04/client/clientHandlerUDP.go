package clientHandlerUDP


import (
	"net"
	"fmt"
)

func create(address string) {
	conn, err := net.Dial("udp", address)
	checkError(err)
	return conn
}

func read(size int, conn net.Conn) []byte {
	buffer := make([]byte, size)
	n, err := conn.Read(buffer)
	checkError(err)
	return buffer
}

func send(buffer []byte, addr net.Addr, conn net.Conn){
	_, err := conn.Write(buffer)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}

