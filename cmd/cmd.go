package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tmuntan1/suse-connect-api-go/config"
	"github.com/tmuntan1/suse-connect-api-go/connect_api"
	"os"
)

var rootCmd = &cobra.Command{}
var apiService connect_api.Service

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	config.Load()
	apiService = connect_api.Service{
		Username: config.Config.Username,
		Password: config.Config.Password,
	}

	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(unscopedCmd)
}
