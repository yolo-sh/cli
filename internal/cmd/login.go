package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/features"
)

// loginCmd represents the "yolo login" command
var loginCmd = &cobra.Command{
	Use: "login",

	Short: "Connect a GitHub account to use with Yolo",

	Long: `Connect a GitHub account to use with Yolo.

Yolo requires the following permissions:

  - "Public SSH keys" and "Repositories" to let you access your repositories from your environments

  - "GPG Keys" and "Personal user data" to configure Git and sign your commits (verified badge)

All your data (including the OAuth access token) will only be stored locally.`,

	Example: "  yolo login",

	Run: func(cmd *cobra.Command, args []string) {
		loginInput := features.LoginInput{}

		login := dependencies.ProvideLoginFeature()

		err := login.Execute(loginInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
