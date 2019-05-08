package node

import (
	"fmt"
	"net"
	"bufio"

	"github.com/mntky/lxd-controller/pkg"
	"github.com/lxc/lxd/shared/api"
)

var raddr = "10.71.173.160"
var rport = ":8080"
var saddr = "10.71.173.139"
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
		go handlerecv(status)
	}
}

func handlerecv(status, containername string) {
	lxdconn := lxdpkg.Connect()
	switch status {
		case "status":
			statresp := lxdpkg.Status(containername, lxdconn)
			send(&statresp)
		case: "create":
			createresp := lxdpkg.Create(containername, lxdconn)
			send(&createresp)
		case: "delete":
			deleteresp := lxdpkg.Delete(containername, lxdconn)
			send(&deleteresp)
		case: "start":
			startresp := lxdpkg.Start(containername, lxdconn)
			send(&startresp)
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
