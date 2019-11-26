package cmd

import (
	"github.com/spf13/cobra"
)

func  NewWoyectl() *cobra.Command {
	cmd := &cobra.Command{
		Use:		"woyectl",
		Short:	"woyectl control lxc",
		Long:		`woyectl control lxc container.
woyectl get node,pod status and spec`,
		Run: runHelp,
	}
	return cmd
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
