package net

import (
	"net"

)

var raddr = "10.16.10.54"
var rport = ":8990"

func Listner() {
	ln, err := net.Listen("tcp", raddr+rport)
	if err != nil {
		return "", err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return "", err
		}
		defer conn.Close()

		var buf bytes.Buffer
		status, err := bufio.NewReader(conn).ReadString("\n")
		err = json.Indent(&buf, []byte(status), "", " ")
		if err != nil {
			return "", err
		}
		indentJson := buf.String()
		fmt.Println(indentJson)
	}
}
