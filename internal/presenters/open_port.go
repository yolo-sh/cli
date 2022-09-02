package presenters

import (
	"fmt"

	"github.com/yolo-sh/cli/internal/constants"
	"github.com/yolo-sh/cli/internal/features"
)

type OpenPortViewDataContent struct {
	ShowAsWarning bool
	Message       string
}

type OpenPortViewData struct {
	Error   *ViewableError
	Content OpenPortViewDataContent
}

type OpenPortViewer interface {
	View(OpenPortViewData)
}

type OpenPortPresenter struct {
	viewableErrorBuilder ViewableErrorBuilder
	viewer               OpenPortViewer
}

func NewOpenPortPresenter(
	viewableErrorBuilder ViewableErrorBuilder,
	viewer OpenPortViewer,
) OpenPortPresenter {

	return OpenPortPresenter{
		viewableErrorBuilder: viewableErrorBuilder,
		viewer:               viewer,
	}
}

func (o OpenPortPresenter) PresentToView(response features.OpenPortResponse) {
	viewData := OpenPortViewData{}

	if response.Error == nil {
		portOpened := response.Content.PortOpened
		envIPAddress := response.Content.EnvPublicIPAddress

		viewDataMessage := fmt.Sprintf(
			"The port \"%s\" is now reachable at: %s",
			portOpened,
			constants.Blue(envIPAddress+":"+portOpened),
		)

		if response.Content.PortAlreadyOpened {
			viewDataMessage = fmt.Sprintf(
				"The port \"%s\" is already open and reachable at: %s",
				portOpened,
				constants.Blue(envIPAddress+":"+portOpened),
			)
		}

		viewData.Content = OpenPortViewDataContent{
			ShowAsWarning: response.Content.PortAlreadyOpened,
			Message:       viewDataMessage,
		}

		o.viewer.View(viewData)

		return
	}

	viewData.Error = o.viewableErrorBuilder.Build(response.Error)

	o.viewer.View(viewData)
}
