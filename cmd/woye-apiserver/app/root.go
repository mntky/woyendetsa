/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package app

import (
  "fmt"

  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var cfgFile string

func NewWoyeapiserver() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "woye-api-server",
		Short: "woye-api-server is LXC cluster frontend",
		RunE: func(cmd *cobra.Command, args []string) error {
			url := viper.GetString("url")
			fmt.Printf("listen on %s \n",url)
			return Run(url)
		},
	}

	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("url", "", "", "api server listen address")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))

	return rootCmd
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  }
	//setconfignameはちゃんと拡張子つけて
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.w8a/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("%s",err))
	}
}
