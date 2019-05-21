package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/mntky/woyendetsa/w8amaster/send"
)

type Act struct {
	Action string
	Option string
	Args  string
}

func NewAct(action, option, conta string) *Act {
	w8aAct := &Act{
		Action: action,
		Option: option,
		Args:  conta,
	}
	return w8aAct
}

func main() {
	app := cli.NewApp()
	app.Name = "w8a"
	app.Usage = "w8a is multihost LXC/LXD manager"
	app.Action = func(c *cli.Context) error {
		if os.Args[1] != "" {
			act := NewAct(os.Args[1], os.Args[2], os.Args[3])
			//change struct -> json
			j, _ := json.Marshal(act)
			//fmt.Println(j)
			str, err := sock.Send(j)
			if err != nil {
				return err
			}
			fmt.Println(str)
		}else {
			fmt.Println("show w8a help `w8a --help` ")
			var err error
			return err
		}
		var err error
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
