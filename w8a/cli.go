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
		if os.Args[1] != "" {
			switch os.Args[1] {
			case "get":
				if (len(os.Args) >= 3) {
					text, err := action.Get(os.Args[2])
					if err != nil{
						return err
					}
					fmt.Println(text)
				}

			case "create":
				if (len(os.Args) >= 3) {
					text, err := action.Create(os.Args[2])
					if err != nil{
						return err
					}
					fmt.Println(text)
				}

			case "delete":
				if (len(os.Args) >= 3) {
					text, err := action.Delete(os.Args[2])
					if err != nil{
						return err
					}
					fmt.Println(text)
				}

			case "describe":
				if (len(os.Args) >= 3) {
					text, err := action.Describe(os.Args[2])
					if err != nil{
						return err
					}
					fmt.Println(text)
				}

			default:
				fmt.Printf("%s option is nothing\n", os.Args[2])
			}
			var err error
			return err
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
