package conn

import (
	"fmt"
	"net"
)

var saddr = "10.71.173.160" //test
var sport = ":8080"

func Send(args []byte) (string, error) {
	fmt.Printf("send :&s", args)
	ln, err := net.Dial("tcp", saddr+sport)
	defer ln.Close()
	if err != nil {
		return "", err
	}

	ln.Write([]byte(args))

	return "ok", nil
}
