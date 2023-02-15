package mysql

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Layout mysql行的布局
func (m *Mysql) Layout(ctx context.Context) *fyne.Container {
	ctl := GetMysqlCtl()
	return container.NewHBox(
		widget.NewLabel("Mysql 状态"),
		ctl.Status(),
		ctl.Start(ctx),
		ctl.Stop(ctx),
	)
}
