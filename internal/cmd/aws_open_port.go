package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/constants"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/yolo/entities"
	"github.com/yolo-sh/yolo/features"
)

// awsOpenPortCmd represents the aws open-port command
var awsOpenPortCmd = &cobra.Command{
	Use: "open-port <repository> <port>",

	Short: "Open a port in an environment",

	Long: `Open a port in a specific environment.

Once a port is opened, it becomes reachable from any IP address using the TCP protocol.`,

	Example: `  yolo aws open-port api 8080
  yolo aws open-port yolo-sh/api 8000`,

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

		awsOpenPortInput := features.OpenPortInput{
			ResolvedRepository: *resolvedRepository,
			PortToOpen:         portToOpen,
		}

		awsOpenPort := dependencies.ProvideAWSOpenPortFeature(
			awsRegion,
			awsProfile,
			awsCredentialsFilePath,
			awsConfigFilePath,
		)

		err = awsOpenPort.Execute(awsOpenPortInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	awsCmd.AddCommand(awsOpenPortCmd)
}
