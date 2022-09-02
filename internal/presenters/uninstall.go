package presenters

import (
	"fmt"

	"github.com/yolo-sh/cli/internal/constants"
	"github.com/yolo-sh/cli/internal/features"
)

type UninstallViewDataContent struct {
	ShowAsWarning bool
	Message       string
	Subtext       string
}

type UninstallViewData struct {
	Error   *ViewableError
	Content UninstallViewDataContent
}

type UninstallViewer interface {
	View(UninstallViewData)
}

type UninstallPresenter struct {
	viewableErrorBuilder ViewableErrorBuilder
	viewer               UninstallViewer
}

func NewUninstallPresenter(
	viewableErrorBuilder ViewableErrorBuilder,
	viewer UninstallViewer,
) UninstallPresenter {

	return UninstallPresenter{
		viewableErrorBuilder: viewableErrorBuilder,
		viewer:               viewer,
	}
}

func (u UninstallPresenter) PresentToView(response features.UninstallResponse) {
	viewData := UninstallViewData{}

	if response.Error == nil {
		bold := constants.Bold

		yoloAlreadyUninstalled := response.Content.YoloAlreadyUninstalled

		viewDataMessage := response.Content.SuccessMessage
		viewDataSubtext := fmt.Sprintf(
			"If you want to remove Yolo entirely:\n\n"+
				"  - Remove the Yolo CLI (located at %s)\n\n"+
				"  - Remove the Yolo configuration (located at %s)\n\n"+
				"  - Unauthorize the Yolo application on GitHub by going to: %s",
			bold(response.Content.YoloExecutablePath),
			bold(response.Content.YoloConfigDirPath),
			bold("https://github.com/settings/applications"),
		)

		if yoloAlreadyUninstalled {
			viewDataMessage = response.Content.AlreadyUninstalledMessage
		}

		viewData.Content = UninstallViewDataContent{
			ShowAsWarning: yoloAlreadyUninstalled,
			Message:       viewDataMessage,
			Subtext:       viewDataSubtext,
		}

		u.viewer.View(viewData)

		return
	}

	viewData.Error = u.viewableErrorBuilder.Build(response.Error)
	u.viewer.View(viewData)
}
