package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/features"
)

// hetznerUninstallCmd represents the hetzner uninstall command
var hetznerUninstallCmd = &cobra.Command{
	Use: "uninstall",

	Short: "Uninstall Yolo from your Hetzner account",

	Long: `Uninstall Yolo from your Hetzner account.

All your environments must be removed before running this command.`,

	Example: "  yolo hetzner uninstall",

	Run: func(cmd *cobra.Command, args []string) {

		hetznerUninstallInput := features.UninstallInput{
			SuccessMessage:            "Yolo has been uninstalled from this region on this Hetzner account.",
			AlreadyUninstalledMessage: "Yolo is already uninstalled in this region on this Hetzner account.",
		}

		hetznerUninstall := dependencies.ProvideHetznerUninstallFeature(
			system.UserConfigDir(),
			hetznerRegion,
			hetznerContext,
		)

		err := hetznerUninstall.Execute(hetznerUninstallInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	hetznerCmd.AddCommand(hetznerUninstallCmd)
}
