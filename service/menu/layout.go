package menu

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// Layout 菜单布局
func (m *Menu) Layout(ctx context.Context) *fyne.Container {
	ctl := GetMenuCtl()
	return container.NewHBox(
		layout.NewSpacer(),
		ctl.Setting(ctx),
		layout.NewSpacer(),
	)
}
