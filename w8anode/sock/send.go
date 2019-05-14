package sock

import (
	//"fmt"
	"net"

)

//node -> master
var daddr = "192.168.11.100"
var dport = ":8090"

func Send(args string) (string, error) {
	ln, err := net.Dial("tcp", daddr+dport)
	defer ln.Close()
	if err != nil {
		return "", err
	}

	ln.Write([]byte(args))
	return "ok", nil
}
