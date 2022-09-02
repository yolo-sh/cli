package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/features"
)

var awsRemoveForceEnvRemove bool

// awsRemoveCmd represents the aws remove command
var awsRemoveCmd = &cobra.Command{
	Use: "remove <repository>",

	Short: "Remove an environment",

	Long: `Remove an existing environment.

The environment will be PERMANENTLY removed along with ALL your edits.
	
There is no going back, so please be sure before running this command.`,

	Example: `  yolo aws remove api
  yolo aws remove yolo-sh/api --force`,

	Args: cobra.ExactArgs(1),

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

		awsRemoveInput := features.RemoveInput{
			ResolvedRepository: *resolvedRepository,
			PreRemoveHook:      dependencies.ProvidePreRemoveHook(),
			ForceRemove:        awsRemoveForceEnvRemove,
			ConfirmRemove: func() (bool, error) {
				logger := system.NewLogger()
				return system.AskForConfirmation(
					logger,
					os.Stdin,
					"All your un-pushed edits will be lost.",
				)
			},
		}

		awsRemove := dependencies.ProvideAWSRemoveFeature(
			awsRegion,
			awsProfile,
			awsCredentialsFilePath,
			awsConfigFilePath,
		)

		err = awsRemove.Execute(awsRemoveInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	awsRemoveCmd.Flags().BoolVar(
		&awsRemoveForceEnvRemove,
		"force",
		false,
		"avoid confirmation",
	)

	awsCmd.AddCommand(awsRemoveCmd)
}
