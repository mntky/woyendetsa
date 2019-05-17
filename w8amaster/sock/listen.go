package main

import (
	"bufio"
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

		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("--sock/listen--")
		fmt.Println(status)

	}
}
