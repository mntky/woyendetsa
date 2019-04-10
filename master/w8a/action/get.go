package action

import (
	"fmt"
	"net"

	//"github.com/lxc/lxd/client"
	//"github.com/lxc/lxd/shared/api"
)

func Get(args string) (string, error){
	saddr := "10.25.10.101"
	sport := ":8989"

	fmt.Println(args)
	ln, err := net.Dial("tcp", saddr+sport)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(ln, args)
	ln.Close()

	return "ok", nil
}
