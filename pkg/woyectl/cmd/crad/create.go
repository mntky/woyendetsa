package cmd

import (
	"github.com/spf13/cobra"
	//"github.com/lxc/lxd/shared/api"
)

var createCmd = &cobra.Command{
	Use:		"create",
	Short:	"create lxc",
	Run:		func(cmd *cobra.Command, args []string) {
		if name, err := cmd.Flags().GetString("name"); err != nil {
			return err
		}
		fmt.Printf("%s ok!!",name)
	},
}

init() {
	RootCmd.AddCommand(createCmd)

	showCmd.Flags().StringP("name", "n" "", "lxc name option")
}
