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

func ProvideAWSClosePortFeature(region, profile, credentialsFilePath, configFilePath string) features.ClosePortFeature {
	return provideAWSClosePortFeature(
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

func provideAWSClosePortFeature(
	userConfigEnvVarsResolverOpts awsProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts awsProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts awsCLI.UserConfigLocalResolverOpts,
) features.ClosePortFeature {
	panic(
		wire.Build(
			viewSet,
			awsServiceBuilderSet,
			awsViewableErrorBuilder,

			stepperSet,

			wire.Bind(new(features.ClosePortOutputHandler), new(featuresCLI.ClosePortOutputHandler)),
			featuresCLI.NewClosePortOutputHandler,

			wire.Bind(new(featuresCLI.ClosePortPresenter), new(presenters.ClosePortPresenter)),
			presenters.NewClosePortPresenter,

			wire.Bind(new(presenters.ClosePortViewer), new(views.ClosePortView)),
			views.NewClosePortView,

			features.NewClosePortFeature,
		),
	)
}
