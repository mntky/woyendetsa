package main

import (
//	"bufio"
//	"bytes"
//	"encoding/json"
	"fmt"
	"net"
)

//node -> master
var laddr = "192.168.11.100"
var lport = ":8090"

func main() {
	ln, err := net.Listen("tcp", laddr+lport)
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn)

/*
		var buf bytes.Buffer
		status, err := bufio.NewReader(conn).ReadString('\n')
		err = json.Indent(&buf, []byte(status), "", " ")
		if err != nil {
			fmt.Println(err)
		}
		indentJson := buf.String()
		fmt.Println(indentJson)
*/
	}
}
