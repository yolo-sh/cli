package presenters

import (
	"fmt"

	"github.com/yolo-sh/cli/internal/constants"
	"github.com/yolo-sh/cli/internal/features"
	"github.com/yolo-sh/cli/internal/globals"
)

type InitViewData struct {
	Error   *ViewableError
	Content InitViewDataContent
}

type InitViewDataContent struct {
	ShowAsWarning bool
	Message       string
	Subtext       string
}

type InitViewer interface {
	View(InitViewData)
}

type InitPresenter struct {
	viewableErrorBuilder ViewableErrorBuilder
	viewer               InitViewer
}

func NewInitPresenter(
	viewableErrorBuilder ViewableErrorBuilder,
	viewer InitViewer,
) InitPresenter {

	return InitPresenter{
		viewableErrorBuilder: viewableErrorBuilder,
		viewer:               viewer,
	}
}

func (i InitPresenter) PresentToView(response features.InitResponse) {
	viewData := InitViewData{}

	if response.Error == nil {
		bold := constants.Bold
		envName := response.Content.EnvName

		viewDataMessage := "The environment for \"" + envName + "\" was initialized."

		envAlreadyCreated := response.Content.EnvAlreadyCreated

		if envAlreadyCreated {
			viewDataMessage = "The environment for \"" + envName + "\" is already initialized."
		}

		currentCloudProviderCmd := string(globals.CurrentCloudProvider)
		currentCloudProviderCmdArgs := globals.CurrentCloudProviderArgs

		if len(currentCloudProviderCmdArgs) > 0 {
			currentCloudProviderCmd += " " + currentCloudProviderCmdArgs
		}

		viewDataSubtext := fmt.Sprintf(
			"The public IP of your environment is: %s\n\n"+
				"To connect to your environment:\n\n"+
				"  - With your editor: `%s`\n\n"+
				"  - With SSH        : `%s`\n\n"+
				"To open a port: `%s`\n\n"+
				"Installed runtimes: %s",
			bold(response.Content.EnvPublicIPAddress),
			bold(constants.Blue("yolo "+currentCloudProviderCmd+" edit "+envName)),
			bold(constants.Blue("ssh "+envName)),
			bold(constants.Blue("yolo "+currentCloudProviderCmd+" open-port "+envName+" <port>")),
			bold(constants.White(
				constants.BGBlue(" docker (latest) ")+" ",
				constants.BGBlue(" docker compose (latest) ")+" ",
				constants.BGBlue(" php (latest) ")+"\n\n",
				constants.BGBlue(" java 17.0 / maven 3.8 ")+" ",
				constants.BGBlue(" node 18.7 (via nvm) ")+" ",
				constants.BGBlue(" python 3.10 (via pyenv) ")+"\n\n",
				constants.BGBlue(" ruby 3.1 (via rvm) ")+" ",
				constants.BGBlue(" rust (latest) ")+" ",
				constants.BGBlue(" go (latest) ")+" ",
			),
			),
		)

		viewData.Content = InitViewDataContent{
			ShowAsWarning: envAlreadyCreated,
			Message:       viewDataMessage,
			Subtext:       viewDataSubtext,
		}

		i.viewer.View(viewData)

		return
	}

	viewData.Error = i.viewableErrorBuilder.Build(response.Error)

	i.viewer.View(viewData)
}
