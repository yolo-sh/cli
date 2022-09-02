// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/yolo-sh/cli/internal/agent"
	featuresCLI "github.com/yolo-sh/cli/internal/features"
	hetznerCLI "github.com/yolo-sh/cli/internal/hetzner"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/views"
	hetznerProviderService "github.com/yolo-sh/hetzner-cloud-provider/service"
	hetznerProviderUserConfig "github.com/yolo-sh/hetzner-cloud-provider/userconfig"
	"github.com/yolo-sh/yolo/features"
)

func ProvideHetznerInitFeature(yoloConfigDir, region, context string) features.InitFeature {
	return provideHetznerInitFeature(
		hetznerProviderUserConfig.EnvVarsResolverOpts{
			Region: region,
		},

		hetznerProviderUserConfig.FilesResolverOpts{
			Region:  region,
			Context: context,
		},

		hetznerCLI.UserConfigLocalResolverOpts{
			Context: context,
		},

		hetznerProviderService.BuilderOpts{
			YoloConfigDir: yoloConfigDir,
		},
	)
}

func provideHetznerInitFeature(
	userConfigEnvVarsResolverOpts hetznerProviderUserConfig.EnvVarsResolverOpts,
	userConfigFilesResolverOpts hetznerProviderUserConfig.FilesResolverOpts,
	userConfigLocalResolverOpts hetznerCLI.UserConfigLocalResolverOpts,
	serviceBuilderOpts hetznerProviderService.BuilderOpts,
) features.InitFeature {
	panic(
		wire.Build(
			viewSet,
			hetznerServiceBuilderSet,
			hetznerViewableErrorBuilder,

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
