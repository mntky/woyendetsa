package node

import (
	"fmt"
	"net"
	"bufio"

)

var raddr = "10.71.173.139"
var rport = ":8080"
var saddr = "10.71.173.160"
var sport = "8080"

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
		//go handlerecv(status, )
	}
}
/*
func handlerecv(status, containername string) {
	fmt.Println(containername)
	switch status {
		case "status":
			fmt.Println("status")
		case "create":
			fmt.Println("create")
		case "delete":
			fmt.Println("delete")
		case "start":
			fmt.Println("start")
		default:
			fmt.Println("other")
	}
	return
}

func send(stat **api.ContainerState) {
	conn, err := net.Dial("tcp", saddr+sport)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(conn, *stat)
}
*/
