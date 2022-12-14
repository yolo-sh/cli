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

func ProvideAWSEditFeature(region, profile, credentialsFilePath, configFilePath string) features.EditFeature {
	return provideAWSEditFeature(
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

func provideAWSEditFeature(
	userConfigEnvVarsResolverOpts awsProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts awsProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts awsCLI.UserConfigLocalResolverOpts,
) features.EditFeature {
	panic(
		wire.Build(
			viewSet,
			awsServiceBuilderSet,
			awsViewableErrorBuilder,

			stepperSet,

			vscodeProcessManagerSet,

			vscodeExtensionsManagerSet,

			wire.Bind(new(features.EditOutputHandler), new(featuresCLI.EditOutputHandler)),
			featuresCLI.NewEditOutputHandler,

			wire.Bind(new(featuresCLI.EditPresenter), new(presenters.EditPresenter)),
			presenters.NewEditPresenter,

			wire.Bind(new(presenters.EditViewer), new(views.EditView)),
			views.NewEditView,

			features.NewEditFeature,
		),
	)
}
