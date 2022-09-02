package views

import "github.com/yolo-sh/cli/internal/presenters"

type OpenPortView struct {
	BaseView
}

func NewOpenPortView(baseView BaseView) OpenPortView {
	return OpenPortView{
		BaseView: baseView,
	}
}

func (o OpenPortView) View(data presenters.OpenPortViewData) {
	if data.Error == nil {
		if data.Content.ShowAsWarning {
			o.ShowWarningView(
				data.Content.Message,
				"",
			)
			return
		}

		o.ShowSuccessView(data.Content.Message, "")
		return
	}

	o.ShowErrorView(data.Error)
}
