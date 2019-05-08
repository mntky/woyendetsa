package conn

import (
	"fmt"
	"net"
)


var saddr = "1.2.3.4"
var sport = ":8008"

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
