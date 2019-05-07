package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"

	"github.com/mntky/woyendetsa/actions"
	"github.com/urfave/cli"
)

var saddr = "10.20.30.40"
var sport = ":8080"

type Act struct {
	Action string
	Option string
	Conta  string
}

func NewAct(action, option, conta string) *Act {
	w8aAct := &Act{
		Action: action,
		Option: option,
		Conta:  conta,
	}
	return w8aAct
}

func main() {
	app := cli.NewApp()
	app.Name = "w8a"
	app.Usage = "w8a is multihost LXC manager"
	app.Action = func(c *cli.Context) error {
		if os.Args[1] != "" {
			act := NewAct(os.Args[1], os.Args[2], os.Args[3])
			//change struct -> json
			j, _ := json.Marshal(act)
			str, err := action.Send(j)
			if err != nil {
				return err
			}
			fmt.Println(str)
		}
		fmt.Println("show w8a help `w8a --help` ")
		var err error
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
