package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

type Args struct {
	Txt string
}

type Str string

func (t *Str) Lower(args *Args, reply *string) error {
	*reply = strings.ToLower(args.Txt)
	return nil
}

func (t *Str) Upper(args *Args, reply *string) error {
	*reply = strings.ToUpper(args.Txt)
	return nil
}

func main() {
	str := new(Str)
	rpc.Register(str)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
