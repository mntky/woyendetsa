package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	//default cobra command
	RootCmd = &cobra.Command{
		Use:		"woyectl",
		Short:	"woyectl control lxc",
		Long:		`woyectl control lxc container.
woyectl get node,pod status and spec`,
		Run: runHelp,
	}

)

// create new cobra command 
func NewWoyectl() *cobra.Command {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("name", "n", "", "lxc name option")
	RootCmd.PersistentFlags().StringP("url", "", "", "api-server url")

	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))

	return RootCmd
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".w8a/config")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.ReadInConfig()
}


func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
