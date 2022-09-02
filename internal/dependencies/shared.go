// go:build wireinject
//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/yolo-sh/cli/internal/config"
	"github.com/yolo-sh/cli/internal/interfaces"
	"github.com/yolo-sh/cli/internal/presenters"
	"github.com/yolo-sh/cli/internal/ssh"
	stepperCLI "github.com/yolo-sh/cli/internal/stepper"
	"github.com/yolo-sh/cli/internal/system"
	"github.com/yolo-sh/cli/internal/views"
	"github.com/yolo-sh/cli/internal/vscode"
	"github.com/yolo-sh/yolo/github"
	"github.com/yolo-sh/yolo/stepper"
)

var viewSet = wire.NewSet(
	wire.Bind(new(views.Displayer), new(system.Displayer)),
	system.NewDisplayer,
	views.NewBaseView,
)

func ProvideBaseView() views.BaseView {
	panic(
		wire.Build(
			viewSet,
		),
	)
}

var yoloViewableErrorBuilder = wire.NewSet(
	wire.Bind(new(presenters.ViewableErrorBuilder), new(presenters.YoloViewableErrorBuilder)),
	presenters.NewYoloViewableErrorBuilder,
)

func ProvideYoloViewableErrorBuilder() presenters.YoloViewableErrorBuilder {
	panic(
		wire.Build(
			yoloViewableErrorBuilder,
		),
	)
}

var githubManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.GitHubManager), new(github.Service)),
	github.NewService,
)

var userConfigManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.UserConfigManager), new(config.UserConfig)),
	config.NewUserConfig,
)

var loggerSet = wire.NewSet(
	wire.Bind(new(interfaces.Logger), new(system.Logger)),
	system.NewLogger,
)

var sshConfigManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.SSHConfigManager), new(ssh.Config)),
	ssh.NewConfigWithDefaultConfigFilePath,
)

var sshKnownHostsManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.SSHKnownHostsManager), new(ssh.KnownHosts)),
	ssh.NewKnownHostsWithDefaultKnownHostsFilePath,
)

var sshKeysManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.SSHKeysManager), new(ssh.Keys)),
	ssh.NewKeysWithDefaultDir,
)

var vscodeProcessManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.VSCodeProcessManager), new(vscode.Process)),
	vscode.NewProcess,
)

var vscodeExtensionsManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.VSCodeExtensionsManager), new(vscode.Extensions)),
	vscode.NewExtensions,
)

var browserManagerSet = wire.NewSet(
	wire.Bind(new(interfaces.BrowserManager), new(system.Browser)),
	system.NewBrowser,
)

var sleeperSet = wire.NewSet(
	wire.Bind(new(interfaces.Sleeper), new(system.Sleeper)),
	system.NewSleeper,
)

var stepperSet = wire.NewSet(
	wire.Bind(new(stepper.Stepper), new(stepperCLI.Stepper)),
	stepperCLI.NewStepper,
)
