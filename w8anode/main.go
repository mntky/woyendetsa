package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/mntky/lxd-controller/pkg"
	"github.com/lxc/lxd/shared/api"
	"github.com/mntky/woyendetsa/w8anode/sock"
)

//master -> node 
var laddr = "192.168.11.100"
var lport = ":8080"
var respstruct *api.ContainerState


type Act struct {
        Action string
        Option string
        Conta  string
}

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

		go handleact(conn)
	}
}


func handleact(data net.Conn) {
	act := new(Act)

	fmt.Printf("--data--\n type: %T \n data: %v \n" ,data,data)

	status, err := bufio.NewReader(data).ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("--status-- \n type: %T \n data: %v \n ---json data---\n ",status,status)

	jsonByte :=([]byte)(status)
	if err := json.Unmarshal(jsonByte, act); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("action: %v \n",act.Action)
	fmt.Printf("options: %v \n",act.Option)
	fmt.Printf("container: %v \n",act.Conta)

	fmt.Println("---lxd connect---")

	lxdconn := lxdpkg.Connect()
	switch act.Action {
		case "status":
			resp := lxdpkg.Status(act.Conta, lxdconn)
			respstruct = resp
			fmt.Println(respstruct.Status)
			sock.Send(respstruct.Status)
		case "create":
			resp, err := lxdpkg.Create(act.Conta, lxdconn)
			if err != nil {
				fmt.Println(err)
			}
			sock.Send(resp)
		case "delete":
			resp, err := lxdpkg.Delete(act.Conta, lxdconn)
			if err != nil {
				fmt.Println(err)
			}
			sock.Send(resp)
		case "start":
			resp, err := lxdpkg.Start(act.Conta, lxdconn)
			if err != nil {
				fmt.Println(err)
			}
			sock.Send(resp)
		default:
			fmt.Println("nothing command")
	}
	return
}
