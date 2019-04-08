package main

import (
	"os"
	"fmt"
	"log"

	"github.com/mntky/woyendetsa/action"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "w8a"
	app.Usage = "w8a is multihost LXC manager"
	app.Action = func(c *cli.Context) error {
		switch os.Args[2] {
		case "get":
			if os.Args[2] != "" {
				text, err := action.Get(os.Args[1], os.Args[2])
				if err != nil{
					return err
				}
				fmt.Println(text)
			}
			fmt.Println("No Argments")
		case "create":
			if os.Args[2] != "" {
				text, err := action.Create(os.Args[1], os.Args[2])
				if err != nil{
					return err
				}
				fmt.Println(text)
			}
			fmt.Println("No Argments")
		case "delete":
			if os.Args[2] != "" {
				text, err := action.Delete(os.Args[1], os.Args[2])
				if err != nil{
					return err
				}
				fmt.Println(text)
			}
			fmt.Println("No Argments")
		case "describe":
			if os.Args[2] != "" {
				text, err := action.Describe(os.Args[1], os.Args[2])
				if err != nil{
					return err
				}
				fmt.Println(text)
			}
			fmt.Println("No Argments")
		default:
			fmt.Printf("%s command is nothing\n", os.Args[1])
		}
		var err error
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
