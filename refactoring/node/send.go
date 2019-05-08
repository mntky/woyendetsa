package net

import (
	"net"

)

var raddr = "10.71.173.160"
var rport = ":8080"
var saddr = "10.71.173.139"
var sport = ":8080"

func Send(stat **api.ContainerState) (string, error) {
	conn, err := net.Dial("tcp", saddr+sport)
	defer conn.Close()
	if err != nil {
		return "", err
	}
	fmt.Fprintf(conn, *stat)
}
