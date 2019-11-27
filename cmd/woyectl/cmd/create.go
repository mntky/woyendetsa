package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	//"github.com/lxc/lxd/shared/api"
)

func init() {
	RootCmd.AddCommand(newcreateCmd())
}

func newcreateCmd() *cobra.Command {
	createCmd := &cobra.Command {
		Use:		"create",
		Short:	"create lxc",
		Run:		func(cmd *cobra.Command, args []string) {
			fmt.Println("dubcommand ok!!")
		},
	}
	return createCmd
}
