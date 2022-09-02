package presenters

import (
	"fmt"

	"github.com/yolo-sh/cli/internal/features"
)

type ClosePortViewDataContent struct {
	ShowAsWarning bool
	Message       string
}

type ClosePortViewData struct {
	Error   *ViewableError
	Content ClosePortViewDataContent
}

type ClosePortViewer interface {
	View(ClosePortViewData)
}

type ClosePortPresenter struct {
	viewableErrorBuilder ViewableErrorBuilder
	viewer               ClosePortViewer
}

func NewClosePortPresenter(
	viewableErrorBuilder ViewableErrorBuilder,
	viewer ClosePortViewer,
) ClosePortPresenter {

	return ClosePortPresenter{
		viewableErrorBuilder: viewableErrorBuilder,
		viewer:               viewer,
	}
}

func (c ClosePortPresenter) PresentToView(response features.ClosePortResponse) {
	viewData := ClosePortViewData{}

	if response.Error == nil {
		portClosed := response.Content.PortClosed

		viewDataMessage := fmt.Sprintf(
			"The port \"%s\" is now closed.",
			portClosed,
		)

		if response.Content.PortAlreadyClosed {
			viewDataMessage = fmt.Sprintf(
				"The port \"%s\" is already closed.",
				portClosed,
			)
		}

		viewData.Content = ClosePortViewDataContent{
			ShowAsWarning: response.Content.PortAlreadyClosed,
			Message:       viewDataMessage,
		}

		c.viewer.View(viewData)

		return
	}

	viewData.Error = c.viewableErrorBuilder.Build(response.Error)

	c.viewer.View(viewData)
}
