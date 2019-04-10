package main

import (
	"fmt"
	"net"
	"bufio"
)


var raddr = "192.168.11.100"
var rport = ":8989"

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
		status, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(status)
	}
}
