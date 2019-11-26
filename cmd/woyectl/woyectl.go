package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"w8a/pkg/woyectl/cmd"
	"w8a/pkg/woyectl/cmd/crad"
)

func main() {
	var log = logrus.New()

	command := cmd.NewWoyectl()
	if err :=command.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
