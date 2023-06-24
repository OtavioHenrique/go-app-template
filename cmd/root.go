package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-app-template",
	Short: "Simple template of golang application",
	Long:  "Any long description can be placed here",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	//any root flag to your app can be used here
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
