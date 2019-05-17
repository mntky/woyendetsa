package main

import (
	"bufio"
//	"bytes"
	"encoding/json"
	"fmt"
	"net"
)

type JsonResp struct {
	Status string
}

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
		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)

		jresp := new(JsonResp)
		jsonByte := ([]byte)(resp)
		if err := json.Unmarshal(jsonByte, jresp); err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Status: %v\n", jresp.Status)

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
