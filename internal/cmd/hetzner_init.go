package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/features"
)

var hetznerInitInstanceType string

// hetznerInitCmd represents the hetzner init command
var hetznerInitCmd = &cobra.Command{
	Use: "init <repository>",

	Short: "Initialize a new environment",

	Long: `Initialize a new environment for a specific GitHub repository.

If the passed repository doesn't contain an account name, your personal account is assumed.`,

	Example: `  yolo hetzner init api
  yolo hetzner init yolo-sh/api --instance-type m4.large`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		yoloViewableErrorBuilder := dependencies.ProvideYoloViewableErrorBuilder()
		baseView := dependencies.ProvideBaseView()

		repository := args[0]
		checkForRepositoryExistence := true

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

		hetznerInitInput := features.InitInput{
			InstanceType:       hetznerInitInstanceType,
			ResolvedRepository: *resolvedRepository,
		}

		hetznerInit := dependencies.ProvideHetznerInitFeature(
			system.UserConfigDir(),
			hetznerRegion,
			hetznerContext,
		)

		err = hetznerInit.Execute(hetznerInitInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	hetznerInitCmd.Flags().StringVar(
		&hetznerInitInstanceType,
		"instance-type",
		"cx11",
		"the instance type used by the environment",
	)

	hetznerCmd.AddCommand(hetznerInitCmd)
}
