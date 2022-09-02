package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/globals"
)

var hetznerContext string
var hetznerRegion string

// hetznerCmd represents the hetzner command
var hetznerCmd = &cobra.Command{
	Use: "hetzner",

	Short: "Use Yolo on Hetzner",

	Long: `Use Yolo on Hetzner.
	
To begin, create your first environment using the command:

  yolo hetzner init <repository>

Once initialized, you may want to connect to it using the command: 

  yolo hetzner edit <repository>

If you don't plan to use this environment again, you could remove it using the command:
	
  yolo hetzner remove <repository>

<repository> may be relative to your personal GitHub account (eg: cli) or fully qualified (eg: my-organization/api).	`,

	Example: `  yolo hetzner init yolo-sh/api --instance-type cx11
  yolo hetzner edit yolo-sh/api
  yolo hetzner remove yolo-sh/api`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		globals.CurrentCloudProvider = globals.HetznerCloudProvider
	},
}

func init() {
	hetznerCmd.Flags().StringVar(
		&hetznerContext,
		"context",
		"",
		"the configuration context to use to access your Hetzner account",
	)

	hetznerCmd.Flags().StringVar(
		&hetznerRegion,
		"region",
		"",
		"the region to use to access your Hetzner account",
	)

	rootCmd.AddCommand(hetznerCmd)
}
