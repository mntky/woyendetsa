package main

import (
	"os"
	"fmt"
	"log"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "w8a"
	app.Usage = "w8a is multihost LXC manager"
	app.Action = func(c *cli.Context) {
		fmt.Printf("%t \n", os.Args)
		fmt.Println(&[]os.Args)
		return 
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
