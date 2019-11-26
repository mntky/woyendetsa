package cmd

import (
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:		"create",
	Short:	"create lxc",
	Run:		func(cmd *cobra.Command, args []string) {
		
	},
}

init() {
	RootCmd.AddCommand(subCmd)
}
