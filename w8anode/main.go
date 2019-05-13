package main

import (
	"bufio"
	//"bytes"
	"encoding/json"
	"fmt"
	"net"

	//"github.com/mntky/lxd-controller/pkg"
)

//master -> node 
var laddr = "192.168.11.100"
var lport = ":8080"

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


/*
	err = json.Indent(&buf, []byte(status), "", " ")
	if err != nil {
		fmt.Println(err)
	}
	indentJson := buf.String()
	fmt.Printf("%T", indentJson)
	fmt.Println(indentJson)


	lxdconn := lxdpkg.Connect()
	switch indentJson.Action {
		case "status":
			statresp := lxdpkg.Status(indentJson.Conta, lxdconn)
			sock.Send(&statresp)
		case "create":
			createresp := lxdpkg.Create(indentJson.Conta, lxdconn)
			sock.Send(&createresp)
		case "delete":
			deleteresp := lxdpkg.Delete(indentJson.Conta, lxdconn)
			sock.Send(&deleteresp)
		case "start":
			startresp := lxdpkg.Start(indentJson.Conta, lxdconn)
			sock.Send(&startresp)
		default:
			fmt.Println("nothing command")
	}
	return
*/
}
