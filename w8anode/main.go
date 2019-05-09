package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"github.com/mntky/lxd-controller/pkg"
)

//master -> node 
var laddr = "1.1.1.1"
var lport = ":8080"

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

		var buf bytes.Buffer
		status, err := bufio.NewReader(conn).ReadString('\n')
		err = json.Indent(&buf, []byte(status), "", " ")
		if err != nil {
			fmt.Println(err)
		}
		indentJson := buf.String()
		fmt.Println(indentJson)
		go handleact(indentJson)
	}
}

func handleact(indentJson string) {
	lxdconn := lxdpkg.Connect()
	switch indentJson {
		case "status":
			statresp := lxdpkg.Status(containername, lxdconn)
			sock.Send(&statresp)
		case "create":
			createresp := lxdpkg.Create(containername, lxdconn)
			sock.Send(&createresp)
		case "delete":
			deleteresp := lxdpkg.Delete(containername, lxdconn)
			sock.Send(&deleteresp)
		case "start":
			startresp := lxdpkg.Start(containername, lxdconn)
			sock.Send(&startresp)
		default:
			fmt.Println("nothing command")
	}
	return
}
