package cmd

import (
	"fmt"
	"os"

	"github.com/khangle880/showroom/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "showroom-cli",
	Short: "Show room CLI tool to get information about live streams",
}

func Execute() {
	rootCmd.Version = utils.AppVersion

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		utils.ClearScreen()
		utils.PrintHeader(utils.AppName, utils.AppDescription)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
