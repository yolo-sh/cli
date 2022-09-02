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

// hetznerOpenPortCmd represents the hetzner open-port command
var hetznerOpenPortCmd = &cobra.Command{
	Use: "open-port <repository> <port>",

	Short: "Open a port in an environment",

	Long: `Open a port in a specific environment.

Once a port is opened, it becomes reachable from any IP address using the TCP protocol.`,

	Example: `  yolo hetzner open-port api 8080
  yolo hetzner open-port yolo-sh/api 8000`,

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

		portToOpen := args[1]

		if err = entities.CheckPortValidity(portToOpen, constants.ReservedPorts); err != nil {
			baseView.ShowErrorViewWithStartingNewLine(
				yoloViewableErrorBuilder.Build(
					err,
				),
			)

			os.Exit(1)
		}

		hetznerOpenPortInput := features.OpenPortInput{
			ResolvedRepository: *resolvedRepository,
			PortToOpen:         portToOpen,
		}

		hetznerOpenPort := dependencies.ProvideHetznerOpenPortFeature(
			system.UserConfigDir(),
			hetznerRegion,
			hetznerContext,
		)

		err = hetznerOpenPort.Execute(hetznerOpenPortInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	hetznerCmd.AddCommand(hetznerOpenPortCmd)
}
