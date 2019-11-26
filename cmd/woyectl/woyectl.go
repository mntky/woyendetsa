package main

import (
	"w8s/cmd/lxc-manager/"
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

	cmd := app.NewLXCManager()
	cmd.Execute()
}
