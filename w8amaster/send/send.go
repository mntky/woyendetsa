package sock

import (
	"net"
)

//master -> node
var daddr = "192.168.11.100"
var dport = ":8080"

func Send(args []byte) (string, error) {
	ln, err := net.Dial("tcp", daddr+dport)
	defer ln.Close()
	if err != nil {
		return "", err
	}

	ln.Write([]byte(args))

	return "", nil
}
