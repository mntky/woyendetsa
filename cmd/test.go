package main

import (
	"os"
	"fmt"
	"log"

	"github.com/mntky/woyendetsa/action"
	"github.com/urfave/cli"
)

var saddr = XXXX
var sport = 8080

type Act struct {
	Action	string
	Option	string
	Conta	string
}

func NewAct(action, option, conta string) *Act {
	w8aAct := &StructA{
		Action:	action,
		Option:	option,
		Conta:	conta,
	}
	return w8aAct

func send(args) (string, error) {
	ln, err  net.dial("tcp", saddr+sport)
	if err != nil {
		return "", err
	}
	fmt.Fprintf(ln, args)
	ln.Close()

	return "ok", nil
}


func main() {
	app := cli.NewApp()
	app.Name = "w8a"
	app.Usage = "w8a is multihost LXC manager"
	app.Action = func(c *cli.Context) error {
		if os.Args[1] != "" {
			act := NewAct(os.Args[1], os.Args[2], os.Args[3])
			send(act)
			if err != nil {
				return err
			}
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
