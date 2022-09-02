package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/features"
)

var hetznerRemoveForceEnvRemove bool

// hetznerRemoveCmd represents the hetzner remove command
var hetznerRemoveCmd = &cobra.Command{
	Use: "remove <repository>",

	Short: "Remove an environment",

	Long: `Remove an existing environment.

The environment will be PERMANENTLY removed along with ALL your edits.
	
There is no going back, so please be sure before running this command.`,

	Example: `  yolo hetzner remove api
  yolo hetzner remove yolo-sh/api --force`,

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

		hetznerRemoveInput := features.RemoveInput{
			ResolvedRepository: *resolvedRepository,
			PreRemoveHook:      dependencies.ProvidePreRemoveHook(),
			ForceRemove:        hetznerRemoveForceEnvRemove,
			ConfirmRemove: func() (bool, error) {
				logger := system.NewLogger()
				return system.AskForConfirmation(
					logger,
					os.Stdin,
					"All your un-pushed edits will be lost.",
				)
			},
		}

		hetznerRemove := dependencies.ProvideHetznerRemoveFeature(
			system.UserConfigDir(),
			hetznerRegion,
			hetznerContext,
		)

		err = hetznerRemove.Execute(hetznerRemoveInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	hetznerRemoveCmd.Flags().BoolVar(
		&hetznerRemoveForceEnvRemove,
		"force",
		false,
		"avoid confirmation",
	)

	hetznerCmd.AddCommand(hetznerRemoveCmd)
}
