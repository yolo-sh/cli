package views

import "github.com/yolo-sh/cli/internal/presenters"

type EditView struct {
	BaseView
}

func NewEditView(baseView BaseView) EditView {
	return EditView{
		BaseView: baseView,
	}
}

func (e EditView) View(data presenters.EditViewData) {
	if data.Error == nil {
		e.ShowSuccessView(data.Content.Message, "")
		return
	}

	e.ShowErrorView(data.Error)
}
