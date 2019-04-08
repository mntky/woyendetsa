package action

import (
	"fmt"

	//"github.com/lxc/lxd/client"
	//"github.com/lxc/lxd/shared/api"
)

func Delete(args1, args2 string) (string, error){
	fmt.Println(args1)
	fmt.Println(args2)

	resp := "ok"
	var err error

	return resp, err
}
