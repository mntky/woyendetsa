package sock

import (
	"fmt"
	"net"
	"encoding/json"

)

type SendJson struct {
	Status string
}

func NewSendJson(status string) *SendJson {
	respjson := &SendJson {
		Status: status,
	}
	return respjson
}

//node -> master
var daddr = "192.168.11.100"
var dport = ":8090"

func Send(args string) (string, error) {
	ln, err := net.Dial("tcp", daddr+dport)
	defer ln.Close()
	if err != nil {
		return "", err
	}
	resp := NewSendJson(args)
	j, _ := json.Marshal(resp)
	if err != nil {
		return "", err
	}

	fmt.Printf("send-test\n%v", j)
	ln.Write([]byte(j))
	return "ok", nil
}
