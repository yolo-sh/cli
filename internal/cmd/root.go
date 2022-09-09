package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yolo-sh/cli/internal/config"
	"github.com/yolo-sh/cli/internal/dependencies"
	"github.com/yolo-sh/cli/internal/exceptions"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/yolo/github"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "yolo",

	Short: "Live environments for any repository in any cloud provider",

	Long: `Yolo - Live environments for any repository in any cloud provider

To begin, run the command "yolo login" to connect your GitHub account.	

From there, the most common workflow is:

  - yolo <cloud_provider> init <repository>   : to initialize an environment for a specific GitHub repository

  - yolo <cloud_provider> edit <repository>   : to connect your preferred editor to an environment

  - yolo <cloud_provider> remove <repository> : to remove an unused environment
	
<repository> may be relative to your personal GitHub account (eg: cli) or fully qualified (eg: my-organization/api).`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		ensureUserIsLoggedIn(cmd)
	},

	TraverseChildren: true,

	Version: "v0.0.9",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(
		ensureYoloCLIRequirements,
		initializeYoloCLIConfig,
		ensureGitHubAccessTokenValidity,
	)
}

func ensureYoloCLIRequirements() {
	missingRequirements := []string{}

	sshCommand := "ssh"
	_, err := exec.LookPath(sshCommand)

	if err != nil {
		missingRequirements = append(
			missingRequirements,
			fmt.Sprintf(
				"OpenSSH client (looked for an \"%s\" command available)",
				sshCommand,
			),
		)
	}

	if len(missingRequirements) > 0 {
		yoloViewableErrorBuilder := dependencies.ProvideYoloViewableErrorBuilder()
		baseView := dependencies.ProvideBaseView()

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
}

func initializeYoloCLIConfig() {
	configDir := system.UserConfigDir()
	configDirPerms := fs.FileMode(0700)

	// Ensure configuration dir exists
	err := os.MkdirAll(
		configDir,
		configDirPerms,
	)
	cobra.CheckErr(err)

	configFilePath := system.UserConfigFilePath()
	configFilePerms := fs.FileMode(0600)

	// Ensure configuration file exists
	f, err := os.OpenFile(
		configFilePath,
		os.O_CREATE,
		configFilePerms,
	)
	cobra.CheckErr(err)
	defer f.Close()

	viper.SetConfigFile(configFilePath)
	cobra.CheckErr(viper.ReadInConfig())
}

// ensureGitHubAccessTokenValidity ensures that
// the github access token has not been
// revoked by user
func ensureGitHubAccessTokenValidity() {
	userConfig := config.NewUserConfig()
	userIsLoggedIn := userConfig.GetBool(config.UserConfigKeyUserIsLoggedIn)

	if !userIsLoggedIn {
		return
	}

	gitHubService := github.NewService()

	githubUser, err := gitHubService.GetAuthenticatedUser(
		userConfig.GetString(
			config.UserConfigKeyGitHubAccessToken,
		),
	)

	if err != nil &&
		gitHubService.IsInvalidAccessTokenError(err) { // User has revoked access token

		userIsLoggedIn = false

		userConfig.Set(
			config.UserConfigKeyUserIsLoggedIn,
			userIsLoggedIn,
		)

		// Error is swallowed here to
		// not confuse user with unexpected error
		_ = userConfig.WriteConfig()
	}

	if err == nil {
		// Update config with updated values from GitHub
		userConfig.PopulateFromGitHubUser(githubUser)

		// Error is swallowed here to
		// not confuse user with unexpected error
		_ = userConfig.WriteConfig()
	}
}

func ensureUserIsLoggedIn(cmd *cobra.Command) {
	userConfig := config.NewUserConfig()
	userIsLoggedIn := userConfig.GetBool(config.UserConfigKeyUserIsLoggedIn)

	if !userIsLoggedIn && cmd != loginCmd {
		yoloViewableErrorBuilder := dependencies.ProvideYoloViewableErrorBuilder()
		baseView := dependencies.ProvideBaseView()

		baseView.ShowErrorViewWithStartingNewLine(
			yoloViewableErrorBuilder.Build(
				exceptions.ErrUserNotLoggedIn,
			),
		)

		os.Exit(1)
	}
}
