package cmd

import (
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/spf13/cobra"
	"github.com/yolo-sh/cli/internal/globals"
)

var awsProfile string
var awsRegion string

var awsCredentialsFilePath string
var awsConfigFilePath string

// awsCmd represents the aws command
var awsCmd = &cobra.Command{
	Use: "aws",

	Short: "Use Yolo on Amazon Web Services",

	Long: `Use Yolo on Amazon Web Services.
	
To begin, create your first environment using the command:

  yolo aws init <repository>

Once initialized, you may want to connect to it using the command: 

  yolo aws edit <repository>

If you don't plan to use this environment again, you could remove it using the command:
	
  yolo aws remove <repository>

<repository> may be relative to your personal GitHub account (eg: cli) or fully qualified (eg: my-organization/api).	`,

	Example: `  yolo aws init yolo-sh/api --instance-type m4.large 
  yolo aws edit yolo-sh/api
  yolo aws remove yolo-sh/api`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		globals.CurrentCloudProvider = globals.AWSCloudProvider
	},
}

func init() {
	awsCmd.Flags().StringVar(
		&awsProfile,
		"profile",
		"",
		"the configuration profile to use to access your AWS account",
	)

	awsCmd.Flags().StringVar(
		&awsRegion,
		"region",
		"",
		"the region to use to access your AWS account",
	)

	awsCredentialsFilePath = config.DefaultSharedCredentialsFilename()
	awsConfigFilePath = config.DefaultSharedConfigFilename()

	rootCmd.AddCommand(awsCmd)
}
