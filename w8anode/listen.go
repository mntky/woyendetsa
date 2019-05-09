package sock

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

//master -> node 
var laddr = "1.1.1.1"
var lport = ":8080"

func Listen() (string, error) {
	ln, err := net.Listen("tcp", laddr+lport)
	if err != nil {
		return "", err
	}

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			return "", err
		}

		var buf bytes.Buffer
		status, err := bufio.NewReader(conn).ReadString('\n')
		err = json.Indent(&buf, []byte(status), "", " ")
		if err != nil {
			return "", err
		}
		indentJson := buf.String()
		fmt.Println(indentJson)
	}
}
