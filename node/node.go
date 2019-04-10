package node

import (
	"fmt"
	"net"
	"bufio"

	"github.com/mntky/lxd-controller/pkg"
	"github.com/lxc/lxd/shared/api"
)

var raddr = "10.25.10.101"
var rport = ":8989"
var saddr = "10.25.10.132"
var sport = "8989"

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
		go handlerecv(status)
	}
}

func handlerecv(status string) {
	switch status {
		case "get":
			lxdconn := lxdpkg.Connect()
			containerstat := lxdpkg.Status("ubuko", lxdconn)
			send(&containerstat)
		default:
			fmt.Println("test")
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
