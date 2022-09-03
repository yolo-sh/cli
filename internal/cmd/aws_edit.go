package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/exceptions"
	"github.com/yolo-sh/cli/internal/vscode"
	"github.com/yolo-sh/yolo/features"
)

// awsEditCmd represents the aws edit command
var awsEditCmd = &cobra.Command{
	Use: "edit <repository>",

	Short: "Connect your editor to an environment",

	Long: `Connect your preferred editor to an environment.

In this version of the Yolo CLI, only Visual Studio Code is supported.`,

	Example: `  yolo aws edit api
  yolo aws edit yolo-sh/api`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		yoloViewableErrorBuilder := dependencies.ProvideYoloViewableErrorBuilder()
		baseView := dependencies.ProvideBaseView()

		missingRequirements := []string{}
		vscodeCLI := vscode.CLI{}
		_, err := vscodeCLI.LookupPath(runtime.GOOS)

		if vscodeCLINotFoundErr, ok := err.(vscode.ErrCLINotFound); ok {
			missingRequirements = append(
				missingRequirements,
				fmt.Sprintf(
					"Visual Studio Code (looked in \"%s)",
					strings.Join(vscodeCLINotFoundErr.VisitedPaths, "\", \"")+"\"",
				),
			)
		}

		if len(missingRequirements) > 0 {
			missingRequirementsErr := exceptions.ErrMissingRequirements{
				MissingRequirements: missingRequirements,
			}

			baseView.ShowErrorViewWithStartingNewLine(
				yoloViewableErrorBuilder.Build(
					missingRequirementsErr,
				),
			)

			os.Exit(1)
		}

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

		awsEditInput := features.EditInput{
			ResolvedRepository: *resolvedRepository,
		}

		awsEdit := dependencies.ProvideAWSEditFeature(
			awsRegion,
			awsProfile,
			awsCredentialsFilePath,
			awsConfigFilePath,
		)

		err = awsEdit.Execute(awsEditInput)

		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	awsCmd.AddCommand(awsEditCmd)
}
