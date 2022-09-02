// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	awsProviderUserConfig "github.com/yolo-sh/aws-cloud-provider/userconfig"
	"github.com/yolo-sh/cli/internal/agent"
	awsCLI "github.com/yolo-sh/cli/internal/aws"
	featuresCLI "github.com/yolo-sh/cli/internal/features"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/views"
	"github.com/yolo-sh/yolo/features"
)

func ProvideAWSInitFeature(region, profile, credentialsFilePath, configFilePath string) features.InitFeature {
	return provideAWSInitFeature(
		awsProviderUserConfig.EnvVarsResolverOpts{
			Region: region,
		},

		awsProviderUserConfig.FilesResolverOpts{
			Region:              region,
			Profile:             profile,
			CredentialsFilePath: credentialsFilePath,
			ConfigFilePath:      configFilePath,
		},

		awsCLI.UserConfigLocalResolverOpts{
			Profile: profile,
		},
	)
}

func provideAWSInitFeature(
	userConfigEnvVarsResolverOpts awsProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts awsProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts awsCLI.UserConfigLocalResolverOpts,
) features.InitFeature {
	panic(
		wire.Build(
			viewSet,
			awsServiceBuilderSet,
			awsViewableErrorBuilder,

			userConfigManagerSet,

			wire.Bind(new(agent.ClientBuilder), new(agent.DefaultClientBuilder)),
			agent.NewDefaultClientBuilder,

			githubManagerSet,

			loggerSet,

			sshConfigManagerSet,

			sshKnownHostsManagerSet,

			sshKeysManagerSet,

			stepperSet,

			wire.Bind(new(features.InitOutputHandler), new(featuresCLI.InitOutputHandler)),
			featuresCLI.NewInitOutputHandler,

			wire.Bind(new(featuresCLI.InitPresenter), new(presenters.InitPresenter)),
			presenters.NewInitPresenter,

			wire.Bind(new(presenters.InitViewer), new(views.InitView)),
			views.NewInitView,

			features.NewInitFeature,
		),
	)
}
