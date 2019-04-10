package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"bytes"
)


var raddr = "10.25.10.113"
var rport = ":8990"

func main() {
	ln, err := net.Listen("tcp", raddr+rport)
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		var buf bytes.Buffer
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("%T\n", status)
		err = json.Indent(&buf, []byte(status), "","  ")
		if err != nil {
			panic(err)
		}
		indentJson := buf.String()
		fmt.Println(indentJson)
		conn.Close()
	}
}
