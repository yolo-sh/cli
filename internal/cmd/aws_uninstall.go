package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/yolo/features"
)

// awsUninstallCmd represents the aws uninstall command
var awsUninstallCmd = &cobra.Command{
	Use: "uninstall",

	Short: "Uninstall Yolo from your AWS account",

	Long: `Uninstall Yolo from your AWS account.

All your environments must be removed before running this command.`,

	Example: "  yolo aws uninstall",

	Run: func(cmd *cobra.Command, args []string) {

		awsUninstallInput := features.UninstallInput{
			SuccessMessage:            "Yolo has been uninstalled from this region on this AWS account.",
			AlreadyUninstalledMessage: "Yolo is already uninstalled in this region on this AWS account.",
		}

		awsUninstall := dependencies.ProvideAWSUninstallFeature(
			awsRegion,
			awsProfile,
			awsCredentialsFilePath,
			awsConfigFilePath,
		)

		err := awsUninstall.Execute(awsUninstallInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	awsCmd.AddCommand(awsUninstallCmd)
}
