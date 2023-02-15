package nginx

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Layout nginx行的布局
func (n *Nginx) Layout(ctx context.Context) *fyne.Container {
	return container.NewHBox(
		widget.NewLabel("Nginx 状态"), // 列表文字
		GetNginxCtl().Status(),      // 服务状态
		GetNginxCtl().Start(ctx),    // 开始按钮
		GetNginxCtl().Stop(ctx),     // 停止按钮
		GetNginxCtl().Restart(ctx),  // 重启按钮
	)
}
