package clientRequestHandler

import (
	"net"
	"clientHandlerTCP"
	"clientHandlerUDP"
)

conn net.Conn
addr net.Addr

func connect(tp string, address string){
 switch  tp{
 case "tcp":
	conn = clientHandlerTCP.create(address)
 case "udp":
	conn = clientHandlerUDP.create(address)
 }
}


func read(int size){
    switch  tp{
	case "tcp":
	   conn = clientHandlerTCP.read(size, conn)
	case "udp":
	   conn, addr = clientHandlerUDP.read(size, conn)
	}
}

func send(tp string, buff []byte){
	switch  tp{
	case "tcp":
	   conn = clientHandlerTCP.send(buff, conn)
	case "udp":
	   conn = clientHandlerUDP.send(buff, addr, conn)
	}
}

func close(){
	conn.Close()
}