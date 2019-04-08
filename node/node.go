package node

import (
	"fmt"
	"net"
	"bufio"

	"github.com/lxc/lxd/shared/api"
)

func main() {
	raddr := "1.1.1.1"
	rport := ":8989"
	saddr := "2.2.2.2"
	sport := "8989"

	ln, err := net.Listen("tcp", raddr+rport)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		status, err := bufio,NewReader(conn).ReadString('/n')
		go handlerecv(status)
	}

func handlerecv(status string) {
	switch status {
		case "get":
			//-------------------
			container, err := lxd.ConnectLXDUnix("", nil)
			if err !=  nil {
				fmt.Println(err)
			}
			//------------------
			stat, str, err := container.GetContainerState(ubuko)
			if err != nil {
				fmt.Println(str)
				fmt.Println(err)
			}
			send(*stat)
		default:
			fmt.Println("test")
	}
	return nil
}

func send(stat) {
	conn, err := net.Dial("tcp", saddr+sport)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(conn, stat)
}
