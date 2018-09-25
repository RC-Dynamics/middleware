package server

import "strings"

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
