package views

import "github.com/yolo-sh/cli/internal/presenters"

type ClosePortView struct {
	BaseView
}

func NewClosePortView(baseView BaseView) ClosePortView {
	return ClosePortView{
		BaseView: baseView,
	}
}

func (c ClosePortView) View(data presenters.ClosePortViewData) {
	if data.Error == nil {
		if data.Content.ShowAsWarning {
			c.ShowWarningView(
				data.Content.Message,
				"",
			)
			return
		}

		c.ShowSuccessView(data.Content.Message, "")
		return
	}

	c.ShowErrorView(data.Error)
}
