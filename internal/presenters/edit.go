package presenters

import (
	"github.com/yolo-sh/cli/internal/features"
)

type EditViewDataContent struct {
	Message string
}

type EditViewData struct {
	Error   *ViewableError
	Content EditViewDataContent
}

type EditViewer interface {
	View(EditViewData)
}

type EditPresenter struct {
	viewableErrorBuilder ViewableErrorBuilder
	viewer               EditViewer
}

func NewEditPresenter(
	viewableErrorBuilder ViewableErrorBuilder,
	viewer EditViewer,
) EditPresenter {

	return EditPresenter{
		viewableErrorBuilder: viewableErrorBuilder,
		viewer:               viewer,
	}
}

func (e EditPresenter) PresentToView(response features.EditResponse) {
	viewData := EditViewData{}

	if response.Error == nil {
		viewDataMessage := "Your editor is now open and connected to your environment."

		viewData.Content = EditViewDataContent{
			Message: viewDataMessage,
		}

		e.viewer.View(viewData)

		return
	}

	viewData.Error = e.viewableErrorBuilder.Build(response.Error)

	e.viewer.View(viewData)
}
