// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/yolo-sh/aws-cloud-provider/config"
	"github.com/yolo-sh/aws-cloud-provider/service"
	"github.com/yolo-sh/aws-cloud-provider/userconfig"
	"github.com/yolo-sh/cli/internal/agent"
	"github.com/yolo-sh/cli/internal/aws"
	config2 "github.com/yolo-sh/cli/internal/config"
	"github.com/yolo-sh/cli/internal/entities"
	features2 "github.com/yolo-sh/cli/internal/features"
	"github.com/yolo-sh/cli/internal/hetzner"
	"github.com/yolo-sh/cli/internal/hooks"
	"github.com/yolo-sh/cli/internal/interfaces"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/ssh"
	"github.com/yolo-sh/cli/internal/stepper"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/cli/internal/views"
	"github.com/yolo-sh/cli/internal/vscode"
	config3 "github.com/yolo-sh/hetzner-cloud-provider/config"
	service2 "github.com/yolo-sh/hetzner-cloud-provider/service"
	userconfig2 "github.com/yolo-sh/hetzner-cloud-provider/userconfig"
	"github.com/yolo-sh/yolo/features"
	"github.com/yolo-sh/yolo/github"
	stepper2 "github.com/yolo-sh/yolo/stepper"
)

// Injectors from aws_close_port.go:

func provideAWSClosePortFeature(userConfigEnvVarsResolverOpts userconfig.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig.FilesResolverOpts, userConfigLocalResolverOpts aws.UserConfigLocalResolverOpts) features.ClosePortFeature {
	stepperStepper := stepper.NewStepper()
	awsAWSViewableErrorBuilder := aws.NewAWSViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	closePortView := views.NewClosePortView(baseView)
	closePortPresenter := presenters.NewClosePortPresenter(awsAWSViewableErrorBuilder, closePortView)
	closePortOutputHandler := features2.NewClosePortOutputHandler(closePortPresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	profileLoader := config.NewProfileLoader()
	filesResolver := userconfig.NewFilesResolver(profileLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := aws.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config.NewUserConfigValidator()
	userConfigLoader := config.NewUserConfigLoader()
	builder := service.NewBuilder(userConfigLocalResolver, userConfigValidator, userConfigLoader)
	closePortFeature := features.NewClosePortFeature(stepperStepper, closePortOutputHandler, builder)
	return closePortFeature
}

// Injectors from aws_edit.go:

func provideAWSEditFeature(userConfigEnvVarsResolverOpts userconfig.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig.FilesResolverOpts, userConfigLocalResolverOpts aws.UserConfigLocalResolverOpts) features.EditFeature {
	stepperStepper := stepper.NewStepper()
	awsAWSViewableErrorBuilder := aws.NewAWSViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	editView := views.NewEditView(baseView)
	editPresenter := presenters.NewEditPresenter(awsAWSViewableErrorBuilder, editView)
	process := vscode.NewProcess()
	extensions := vscode.NewExtensions()
	editOutputHandler := features2.NewEditOutputHandler(editPresenter, process, extensions)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	profileLoader := config.NewProfileLoader()
	filesResolver := userconfig.NewFilesResolver(profileLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := aws.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config.NewUserConfigValidator()
	userConfigLoader := config.NewUserConfigLoader()
	builder := service.NewBuilder(userConfigLocalResolver, userConfigValidator, userConfigLoader)
	editFeature := features.NewEditFeature(stepperStepper, editOutputHandler, builder)
	return editFeature
}

// Injectors from aws_init.go:

func provideAWSInitFeature(userConfigEnvVarsResolverOpts userconfig.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig.FilesResolverOpts, userConfigLocalResolverOpts aws.UserConfigLocalResolverOpts) features.InitFeature {
	stepperStepper := stepper.NewStepper()
	userConfig := config2.NewUserConfig()
	awsAWSViewableErrorBuilder := aws.NewAWSViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	initView := views.NewInitView(baseView)
	initPresenter := presenters.NewInitPresenter(awsAWSViewableErrorBuilder, initView)
	defaultClientBuilder := agent.NewDefaultClientBuilder()
	githubService := github.NewService()
	logger := system.NewLogger()
	sshConfig := ssh.NewConfigWithDefaultConfigFilePath()
	keys := ssh.NewKeysWithDefaultDir()
	knownHosts := ssh.NewKnownHostsWithDefaultKnownHostsFilePath()
	initOutputHandler := features2.NewInitOutputHandler(userConfig, initPresenter, defaultClientBuilder, githubService, logger, sshConfig, keys, knownHosts)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	profileLoader := config.NewProfileLoader()
	filesResolver := userconfig.NewFilesResolver(profileLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := aws.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config.NewUserConfigValidator()
	userConfigLoader := config.NewUserConfigLoader()
	builder := service.NewBuilder(userConfigLocalResolver, userConfigValidator, userConfigLoader)
	initFeature := features.NewInitFeature(stepperStepper, initOutputHandler, builder)
	return initFeature
}

// Injectors from aws_open_port.go:

func provideAWSOpenPortFeature(userConfigEnvVarsResolverOpts userconfig.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig.FilesResolverOpts, userConfigLocalResolverOpts aws.UserConfigLocalResolverOpts) features.OpenPortFeature {
	stepperStepper := stepper.NewStepper()
	awsAWSViewableErrorBuilder := aws.NewAWSViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	openPortView := views.NewOpenPortView(baseView)
	openPortPresenter := presenters.NewOpenPortPresenter(awsAWSViewableErrorBuilder, openPortView)
	openPortOutputHandler := features2.NewOpenPortOutputHandler(openPortPresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	profileLoader := config.NewProfileLoader()
	filesResolver := userconfig.NewFilesResolver(profileLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := aws.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config.NewUserConfigValidator()
	userConfigLoader := config.NewUserConfigLoader()
	builder := service.NewBuilder(userConfigLocalResolver, userConfigValidator, userConfigLoader)
	openPortFeature := features.NewOpenPortFeature(stepperStepper, openPortOutputHandler, builder)
	return openPortFeature
}

// Injectors from aws_remove.go:

func provideAWSRemoveFeature(userConfigEnvVarsResolverOpts userconfig.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig.FilesResolverOpts, userConfigLocalResolverOpts aws.UserConfigLocalResolverOpts) features.RemoveFeature {
	stepperStepper := stepper.NewStepper()
	awsAWSViewableErrorBuilder := aws.NewAWSViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	removeView := views.NewRemoveView(baseView)
	removePresenter := presenters.NewRemovePresenter(awsAWSViewableErrorBuilder, removeView)
	removeOutputHandler := features2.NewRemoveOutputHandler(removePresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	profileLoader := config.NewProfileLoader()
	filesResolver := userconfig.NewFilesResolver(profileLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := aws.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config.NewUserConfigValidator()
	userConfigLoader := config.NewUserConfigLoader()
	builder := service.NewBuilder(userConfigLocalResolver, userConfigValidator, userConfigLoader)
	removeFeature := features.NewRemoveFeature(stepperStepper, removeOutputHandler, builder)
	return removeFeature
}

// Injectors from aws_uninstall.go:

func provideAWSUninstallFeature(userConfigEnvVarsResolverOpts userconfig.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig.FilesResolverOpts, userConfigLocalResolverOpts aws.UserConfigLocalResolverOpts) features.UninstallFeature {
	stepperStepper := stepper.NewStepper()
	awsAWSViewableErrorBuilder := aws.NewAWSViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	uninstallView := views.NewUninstallView(baseView)
	uninstallPresenter := presenters.NewUninstallPresenter(awsAWSViewableErrorBuilder, uninstallView)
	uninstallOutputHandler := features2.NewUninstallOutputHandler(uninstallPresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	profileLoader := config.NewProfileLoader()
	filesResolver := userconfig.NewFilesResolver(profileLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := aws.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config.NewUserConfigValidator()
	userConfigLoader := config.NewUserConfigLoader()
	builder := service.NewBuilder(userConfigLocalResolver, userConfigValidator, userConfigLoader)
	uninstallFeature := features.NewUninstallFeature(stepperStepper, uninstallOutputHandler, builder)
	return uninstallFeature
}

// Injectors from entities.go:

func ProvideEnvRepositoryResolver() entities.EnvRepositoryResolver {
	logger := system.NewLogger()
	userConfig := config2.NewUserConfig()
	githubService := github.NewService()
	envRepositoryResolver := entities.NewEnvRepositoryResolver(logger, userConfig, githubService)
	return envRepositoryResolver
}

// Injectors from hetzner_close_port.go:

func provideHetznerClosePortFeature(userConfigEnvVarsResolverOpts userconfig2.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig2.FilesResolverOpts, userConfigLocalResolverOpts hetzner.UserConfigLocalResolverOpts, serviceBuilderOpts service2.BuilderOpts) features.ClosePortFeature {
	stepperStepper := stepper.NewStepper()
	hetznerHetznerViewableErrorBuilder := hetzner.NewHetznerViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	closePortView := views.NewClosePortView(baseView)
	closePortPresenter := presenters.NewClosePortPresenter(hetznerHetznerViewableErrorBuilder, closePortView)
	closePortOutputHandler := features2.NewClosePortOutputHandler(closePortPresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig2.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	contextLoader := config3.NewContextLoader()
	filesResolver := userconfig2.NewFilesResolver(contextLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := hetzner.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config3.NewUserConfigValidator()
	builder := service2.NewBuilder(serviceBuilderOpts, userConfigLocalResolver, userConfigValidator)
	closePortFeature := features.NewClosePortFeature(stepperStepper, closePortOutputHandler, builder)
	return closePortFeature
}

// Injectors from hetzner_edit.go:

func provideHetznerEditFeature(userConfigEnvVarsResolverOpts userconfig2.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig2.FilesResolverOpts, userConfigLocalResolverOpts hetzner.UserConfigLocalResolverOpts, serviceBuilderOpts service2.BuilderOpts) features.EditFeature {
	stepperStepper := stepper.NewStepper()
	hetznerHetznerViewableErrorBuilder := hetzner.NewHetznerViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	editView := views.NewEditView(baseView)
	editPresenter := presenters.NewEditPresenter(hetznerHetznerViewableErrorBuilder, editView)
	process := vscode.NewProcess()
	extensions := vscode.NewExtensions()
	editOutputHandler := features2.NewEditOutputHandler(editPresenter, process, extensions)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig2.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	contextLoader := config3.NewContextLoader()
	filesResolver := userconfig2.NewFilesResolver(contextLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := hetzner.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config3.NewUserConfigValidator()
	builder := service2.NewBuilder(serviceBuilderOpts, userConfigLocalResolver, userConfigValidator)
	editFeature := features.NewEditFeature(stepperStepper, editOutputHandler, builder)
	return editFeature
}

// Injectors from hetzner_init.go:

func provideHetznerInitFeature(userConfigEnvVarsResolverOpts userconfig2.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig2.FilesResolverOpts, userConfigLocalResolverOpts hetzner.UserConfigLocalResolverOpts, serviceBuilderOpts service2.BuilderOpts) features.InitFeature {
	stepperStepper := stepper.NewStepper()
	userConfig := config2.NewUserConfig()
	hetznerHetznerViewableErrorBuilder := hetzner.NewHetznerViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	initView := views.NewInitView(baseView)
	initPresenter := presenters.NewInitPresenter(hetznerHetznerViewableErrorBuilder, initView)
	defaultClientBuilder := agent.NewDefaultClientBuilder()
	githubService := github.NewService()
	logger := system.NewLogger()
	sshConfig := ssh.NewConfigWithDefaultConfigFilePath()
	keys := ssh.NewKeysWithDefaultDir()
	knownHosts := ssh.NewKnownHostsWithDefaultKnownHostsFilePath()
	initOutputHandler := features2.NewInitOutputHandler(userConfig, initPresenter, defaultClientBuilder, githubService, logger, sshConfig, keys, knownHosts)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig2.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	contextLoader := config3.NewContextLoader()
	filesResolver := userconfig2.NewFilesResolver(contextLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := hetzner.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config3.NewUserConfigValidator()
	builder := service2.NewBuilder(serviceBuilderOpts, userConfigLocalResolver, userConfigValidator)
	initFeature := features.NewInitFeature(stepperStepper, initOutputHandler, builder)
	return initFeature
}

// Injectors from hetzner_open_port.go:

func provideHetznerOpenPortFeature(userConfigEnvVarsResolverOpts userconfig2.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig2.FilesResolverOpts, userConfigLocalResolverOpts hetzner.UserConfigLocalResolverOpts, serviceBuilderOpts service2.BuilderOpts) features.OpenPortFeature {
	stepperStepper := stepper.NewStepper()
	hetznerHetznerViewableErrorBuilder := hetzner.NewHetznerViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	openPortView := views.NewOpenPortView(baseView)
	openPortPresenter := presenters.NewOpenPortPresenter(hetznerHetznerViewableErrorBuilder, openPortView)
	openPortOutputHandler := features2.NewOpenPortOutputHandler(openPortPresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig2.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	contextLoader := config3.NewContextLoader()
	filesResolver := userconfig2.NewFilesResolver(contextLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := hetzner.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config3.NewUserConfigValidator()
	builder := service2.NewBuilder(serviceBuilderOpts, userConfigLocalResolver, userConfigValidator)
	openPortFeature := features.NewOpenPortFeature(stepperStepper, openPortOutputHandler, builder)
	return openPortFeature
}

// Injectors from hetzner_remove.go:

func provideHetznerRemoveFeature(userConfigEnvVarsResolverOpts userconfig2.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig2.FilesResolverOpts, userConfigLocalResolverOpts hetzner.UserConfigLocalResolverOpts, serviceBuilderOpts service2.BuilderOpts) features.RemoveFeature {
	stepperStepper := stepper.NewStepper()
	hetznerHetznerViewableErrorBuilder := hetzner.NewHetznerViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	removeView := views.NewRemoveView(baseView)
	removePresenter := presenters.NewRemovePresenter(hetznerHetznerViewableErrorBuilder, removeView)
	removeOutputHandler := features2.NewRemoveOutputHandler(removePresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig2.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	contextLoader := config3.NewContextLoader()
	filesResolver := userconfig2.NewFilesResolver(contextLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := hetzner.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config3.NewUserConfigValidator()
	builder := service2.NewBuilder(serviceBuilderOpts, userConfigLocalResolver, userConfigValidator)
	removeFeature := features.NewRemoveFeature(stepperStepper, removeOutputHandler, builder)
	return removeFeature
}

// Injectors from hetzner_uninstall.go:

func provideHetznerUninstallFeature(userConfigEnvVarsResolverOpts userconfig2.EnvVarsResolverOpts, userConfigFilesResolverOpts userconfig2.FilesResolverOpts, userConfigLocalResolverOpts hetzner.UserConfigLocalResolverOpts, serviceBuilderOpts service2.BuilderOpts) features.UninstallFeature {
	stepperStepper := stepper.NewStepper()
	hetznerHetznerViewableErrorBuilder := hetzner.NewHetznerViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	uninstallView := views.NewUninstallView(baseView)
	uninstallPresenter := presenters.NewUninstallPresenter(hetznerHetznerViewableErrorBuilder, uninstallView)
	uninstallOutputHandler := features2.NewUninstallOutputHandler(uninstallPresenter)
	envVars := system.NewEnvVars()
	envVarsResolver := userconfig2.NewEnvVarsResolver(envVars, userConfigEnvVarsResolverOpts)
	contextLoader := config3.NewContextLoader()
	filesResolver := userconfig2.NewFilesResolver(contextLoader, userConfigFilesResolverOpts, envVars)
	userConfigLocalResolver := hetzner.NewUserConfigLocalResolver(envVarsResolver, filesResolver, userConfigLocalResolverOpts)
	userConfigValidator := config3.NewUserConfigValidator()
	builder := service2.NewBuilder(serviceBuilderOpts, userConfigLocalResolver, userConfigValidator)
	uninstallFeature := features.NewUninstallFeature(stepperStepper, uninstallOutputHandler, builder)
	return uninstallFeature
}

// Injectors from hooks.go:

func ProvidePreRemoveHook() hooks.PreRemove {
	sshConfig := ssh.NewConfigWithDefaultConfigFilePath()
	keys := ssh.NewKeysWithDefaultDir()
	knownHosts := ssh.NewKnownHostsWithDefaultKnownHostsFilePath()
	userConfig := config2.NewUserConfig()
	githubService := github.NewService()
	preRemove := hooks.NewPreRemove(sshConfig, keys, knownHosts, userConfig, githubService)
	return preRemove
}

// Injectors from login.go:

func ProvideLoginFeature() features2.LoginFeature {
	presentersYoloViewableErrorBuilder := presenters.NewYoloViewableErrorBuilder()
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	loginView := views.NewLoginView(baseView)
	loginPresenter := presenters.NewLoginPresenter(presentersYoloViewableErrorBuilder, loginView)
	logger := system.NewLogger()
	browser := system.NewBrowser()
	userConfig := config2.NewUserConfig()
	sleeper := system.NewSleeper()
	githubService := github.NewService()
	loginFeature := features2.NewLoginFeature(loginPresenter, logger, browser, userConfig, sleeper, githubService)
	return loginFeature
}

// Injectors from shared.go:

func ProvideBaseView() views.BaseView {
	displayer := system.NewDisplayer()
	baseView := views.NewBaseView(displayer)
	return baseView
}

func ProvideYoloViewableErrorBuilder() presenters.YoloViewableErrorBuilder {
	presentersYoloViewableErrorBuilder := presenters.NewYoloViewableErrorBuilder()
	return presentersYoloViewableErrorBuilder
}

// aws_close_port.go:

func ProvideAWSClosePortFeature(region, profile, credentialsFilePath, configFilePath string) features.ClosePortFeature {
	return provideAWSClosePortFeature(userconfig.EnvVarsResolverOpts{
		Region: region,
	}, userconfig.FilesResolverOpts{
		Region:              region,
		Profile:             profile,
		CredentialsFilePath: credentialsFilePath,
		ConfigFilePath:      configFilePath,
	}, aws.UserConfigLocalResolverOpts{
		Profile: profile,
	},
	)
}

// aws_edit.go:

func ProvideAWSEditFeature(region, profile, credentialsFilePath, configFilePath string) features.EditFeature {
	return provideAWSEditFeature(userconfig.EnvVarsResolverOpts{
		Region: region,
	}, userconfig.FilesResolverOpts{
		Region:              region,
		Profile:             profile,
		CredentialsFilePath: credentialsFilePath,
		ConfigFilePath:      configFilePath,
	}, aws.UserConfigLocalResolverOpts{
		Profile: profile,
	},
	)
}

// aws_init.go:

func ProvideAWSInitFeature(region, profile, credentialsFilePath, configFilePath string) features.InitFeature {
	return provideAWSInitFeature(userconfig.EnvVarsResolverOpts{
		Region: region,
	}, userconfig.FilesResolverOpts{
		Region:              region,
		Profile:             profile,
		CredentialsFilePath: credentialsFilePath,
		ConfigFilePath:      configFilePath,
	}, aws.UserConfigLocalResolverOpts{
		Profile: profile,
	},
	)
}

// aws_open_port.go:

func ProvideAWSOpenPortFeature(region, profile, credentialsFilePath, configFilePath string) features.OpenPortFeature {
	return provideAWSOpenPortFeature(userconfig.EnvVarsResolverOpts{
		Region: region,
	}, userconfig.FilesResolverOpts{
		Region:              region,
		Profile:             profile,
		CredentialsFilePath: credentialsFilePath,
		ConfigFilePath:      configFilePath,
	}, aws.UserConfigLocalResolverOpts{
		Profile: profile,
	},
	)
}

// aws_remove.go:

func ProvideAWSRemoveFeature(region, profile, credentialsFilePath, configFilePath string) features.RemoveFeature {
	return provideAWSRemoveFeature(userconfig.EnvVarsResolverOpts{
		Region: region,
	}, userconfig.FilesResolverOpts{
		Region:              region,
		Profile:             profile,
		CredentialsFilePath: credentialsFilePath,
		ConfigFilePath:      configFilePath,
	}, aws.UserConfigLocalResolverOpts{
		Profile: profile,
	},
	)
}

// aws_uninstall.go:

func ProvideAWSUninstallFeature(region, profile, credentialsFilePath, configFilePath string) features.UninstallFeature {
	return provideAWSUninstallFeature(userconfig.EnvVarsResolverOpts{
		Region: region,
	}, userconfig.FilesResolverOpts{
		Region:              region,
		Profile:             profile,
		CredentialsFilePath: credentialsFilePath,
		ConfigFilePath:      configFilePath,
	}, aws.UserConfigLocalResolverOpts{
		Profile: profile,
	},
	)
}

// hetzner_close_port.go:

func ProvideHetznerClosePortFeature(yoloConfigDir, region, context string) features.ClosePortFeature {
	return provideHetznerClosePortFeature(userconfig2.EnvVarsResolverOpts{
		Region: region,
	}, userconfig2.FilesResolverOpts{
		Region:  region,
		Context: context,
	}, hetzner.UserConfigLocalResolverOpts{
		Context: context,
	}, service2.BuilderOpts{
		YoloConfigDir: yoloConfigDir,
	},
	)
}

// hetzner_edit.go:

func ProvideHetznerEditFeature(yoloConfigDir, region, context string) features.EditFeature {
	return provideHetznerEditFeature(userconfig2.EnvVarsResolverOpts{
		Region: region,
	}, userconfig2.FilesResolverOpts{
		Region:  region,
		Context: context,
	}, hetzner.UserConfigLocalResolverOpts{
		Context: context,
	}, service2.BuilderOpts{
		YoloConfigDir: yoloConfigDir,
	},
	)
}

// hetzner_init.go:

func ProvideHetznerInitFeature(yoloConfigDir, region, context string) features.InitFeature {
	return provideHetznerInitFeature(userconfig2.EnvVarsResolverOpts{
		Region: region,
	}, userconfig2.FilesResolverOpts{
		Region:  region,
		Context: context,
	}, hetzner.UserConfigLocalResolverOpts{
		Context: context,
	}, service2.BuilderOpts{
		YoloConfigDir: yoloConfigDir,
	},
	)
}

// hetzner_open_port.go:

func ProvideHetznerOpenPortFeature(yoloConfigDir, region, context string) features.OpenPortFeature {
	return provideHetznerOpenPortFeature(userconfig2.EnvVarsResolverOpts{
		Region: region,
	}, userconfig2.FilesResolverOpts{
		Region:  region,
		Context: context,
	}, hetzner.UserConfigLocalResolverOpts{
		Context: context,
	}, service2.BuilderOpts{
		YoloConfigDir: yoloConfigDir,
	},
	)
}

// hetzner_remove.go:

func ProvideHetznerRemoveFeature(yoloConfigDir, region, context string) features.RemoveFeature {
	return provideHetznerRemoveFeature(userconfig2.EnvVarsResolverOpts{
		Region: region,
	}, userconfig2.FilesResolverOpts{
		Region:  region,
		Context: context,
	}, hetzner.UserConfigLocalResolverOpts{
		Context: context,
	}, service2.BuilderOpts{
		YoloConfigDir: yoloConfigDir,
	},
	)
}

// hetzner_uninstall.go:

func ProvideHetznerUninstallFeature(yoloConfigDir, region, context string) features.UninstallFeature {
	return provideHetznerUninstallFeature(userconfig2.EnvVarsResolverOpts{
		Region: region,
	}, userconfig2.FilesResolverOpts{
		Region:  region,
		Context: context,
	}, hetzner.UserConfigLocalResolverOpts{
		Context: context,
	}, service2.BuilderOpts{
		YoloConfigDir: yoloConfigDir,
	},
	)
}

// shared.go:

var viewSet = wire.NewSet(wire.Bind(new(views.Displayer), new(system.Displayer)), system.NewDisplayer, views.NewBaseView)

var yoloViewableErrorBuilder = wire.NewSet(wire.Bind(new(presenters.ViewableErrorBuilder), new(presenters.YoloViewableErrorBuilder)), presenters.NewYoloViewableErrorBuilder)

var githubManagerSet = wire.NewSet(wire.Bind(new(interfaces.GitHubManager), new(github.Service)), github.NewService)

var userConfigManagerSet = wire.NewSet(wire.Bind(new(interfaces.UserConfigManager), new(config2.UserConfig)), config2.NewUserConfig)

var loggerSet = wire.NewSet(wire.Bind(new(interfaces.Logger), new(system.Logger)), system.NewLogger)

var sshConfigManagerSet = wire.NewSet(wire.Bind(new(interfaces.SSHConfigManager), new(ssh.Config)), ssh.NewConfigWithDefaultConfigFilePath)

var sshKnownHostsManagerSet = wire.NewSet(wire.Bind(new(interfaces.SSHKnownHostsManager), new(ssh.KnownHosts)), ssh.NewKnownHostsWithDefaultKnownHostsFilePath)

var sshKeysManagerSet = wire.NewSet(wire.Bind(new(interfaces.SSHKeysManager), new(ssh.Keys)), ssh.NewKeysWithDefaultDir)

var vscodeProcessManagerSet = wire.NewSet(wire.Bind(new(interfaces.VSCodeProcessManager), new(vscode.Process)), vscode.NewProcess)

var vscodeExtensionsManagerSet = wire.NewSet(wire.Bind(new(interfaces.VSCodeExtensionsManager), new(vscode.Extensions)), vscode.NewExtensions)

var browserManagerSet = wire.NewSet(wire.Bind(new(interfaces.BrowserManager), new(system.Browser)), system.NewBrowser)

var sleeperSet = wire.NewSet(wire.Bind(new(interfaces.Sleeper), new(system.Sleeper)), system.NewSleeper)

var stepperSet = wire.NewSet(wire.Bind(new(stepper2.Stepper), new(stepper.Stepper)), stepper.NewStepper)
