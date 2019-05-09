package sock

import (
	"fmt"
	"net"
)

//master -> node
var daddr = "1.2.3.4"
var dport = ":8008"

func Send(args []byte) (string, error) {
	fmt.Printf("send :&s", args)
	ln, err := net.Dial("tcp", daddr+dport)
	defer ln.Close()
	if err != nil {
		return "", err
	}

	ln.Write([]byte(args))

	return "ok", nil
}
