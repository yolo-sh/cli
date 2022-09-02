package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/constants"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/features"
)

// hetznerClosePortCmd represents the hetzner close-port command
var hetznerClosePortCmd = &cobra.Command{
	Use: "close-port <repository> <port>",

	Short: "Close a port in an environment",

	Long: `Close a port in a specific environment.

Once a port is closed, it becomes unreachable from any IP address.`,

	Example: `  yolo hetzner close-port api 8080
  yolo hetzner close-port yolo-sh/api 8000`,

	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {

		yoloViewableErrorBuilder := dependencies.ProvideYoloViewableErrorBuilder()
		baseView := dependencies.ProvideBaseView()

		repository := args[0]
		checkForRepositoryExistence := false

		repositoryResolver := dependencies.ProvideEnvRepositoryResolver()
		resolvedRepository, err := repositoryResolver.Resolve(
			repository,
			checkForRepositoryExistence,
		)

		if err != nil {
			baseView.ShowErrorViewWithStartingNewLine(
				yoloViewableErrorBuilder.Build(
					err,
				),
			)

			os.Exit(1)
		}

		portToClose := args[1]

		if err = entities.CheckPortValidity(portToClose, constants.ReservedPorts); err != nil {
			baseView.ShowErrorViewWithStartingNewLine(
				yoloViewableErrorBuilder.Build(
					err,
				),
			)

			os.Exit(1)
		}

		hetznerClosePortInput := features.ClosePortInput{
			ResolvedRepository: *resolvedRepository,
			PortToClose:        portToClose,
		}

		hetznerClosePort := dependencies.ProvideHetznerClosePortFeature(
			system.UserConfigDir(),
			hetznerRegion,
			hetznerContext,
		)

		err = hetznerClosePort.Execute(hetznerClosePortInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	hetznerCmd.AddCommand(hetznerClosePortCmd)
}
