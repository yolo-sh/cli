// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	awsProviderUserConfig "github.com/yolo-sh/aws-cloud-provider/userconfig"
	awsCLI "github.com/yolo-sh/cli/internal/aws"
	featuresCLI "github.com/yolo-sh/cli/internal/features"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/views"
	"github.com/yolo-sh/yolo/features"
)

func ProvideAWSOpenPortFeature(region, profile, credentialsFilePath, configFilePath string) features.OpenPortFeature {
	return provideAWSOpenPortFeature(
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

func provideAWSOpenPortFeature(
	userConfigEnvVarsResolverOpts awsProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts awsProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts awsCLI.UserConfigLocalResolverOpts,
) features.OpenPortFeature {
	panic(
		wire.Build(
			viewSet,
			awsServiceBuilderSet,
			awsViewableErrorBuilder,

			stepperSet,

			wire.Bind(new(features.OpenPortOutputHandler), new(featuresCLI.OpenPortOutputHandler)),
			featuresCLI.NewOpenPortOutputHandler,

			wire.Bind(new(featuresCLI.OpenPortPresenter), new(presenters.OpenPortPresenter)),
			presenters.NewOpenPortPresenter,

			wire.Bind(new(presenters.OpenPortViewer), new(views.OpenPortView)),
			views.NewOpenPortView,

			features.NewOpenPortFeature,
		),
	)
}
