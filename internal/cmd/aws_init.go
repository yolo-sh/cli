package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/yolo/features"
)

var awsInitInstanceType string

// awsInitCmd represents the aws init command
var awsInitCmd = &cobra.Command{
	Use: "init <repository>",

	Short: "Initialize a new environment",

	Long: `Initialize a new environment for a specific GitHub repository.

If the passed repository doesn't contain an account name, your personal account is assumed.`,

	Example: `  yolo aws init api
  yolo aws init yolo-sh/api --instance-type m4.large`,

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

		awsInitInput := features.InitInput{
			InstanceType:       awsInitInstanceType,
			ResolvedRepository: *resolvedRepository,
		}

		awsInit := dependencies.ProvideAWSInitFeature(
			awsRegion,
			awsProfile,
			awsCredentialsFilePath,
			awsConfigFilePath,
		)

		err = awsInit.Execute(awsInitInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	awsInitCmd.Flags().StringVar(
		&awsInitInstanceType,
		"instance-type",
		"t2.medium",
		"the instance type used by the environment",
	)

	awsCmd.AddCommand(awsInitCmd)
}
