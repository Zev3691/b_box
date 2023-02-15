package nginx

import (
	"b_box/service/nginx/exec"
	"b_box/util"
	"context"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"sync"
)

type Nginx struct {
	Mux     *sync.RWMutex
	Running bool
	status  *canvas.Text
}

var once sync.Once

// 全局单例
var nginxInstance *Nginx

func init() {
	newNginx()
}

func newNginx() *Nginx {
	return &Nginx{
		Mux:     &sync.RWMutex{},
		Running: false,
		status:  canvas.NewText("停止中", util.Red),
	}
}

func getNginxInstance() *Nginx {
	if nginxInstance == nil {
		once.Do(
			func() {
				nginxInstance = newNginx()
			},
		)
	}

	return nginxInstance
}

// GetNginxCtl 获取全局操作单例
func GetNginxCtl() *Nginx {
	return getNginxInstance()
}

// ChangeStatus 改变当前状态
func (n *Nginx) ChangeStatus(nextStatus string, nextStatusColor color.Color) {
	n.status.Text = nextStatus
	n.status.Color = nextStatusColor
	n.status.Refresh()
	return
}

// Start 启动nginx
func (n *Nginx) Start(ctx context.Context) *widget.Button {
	return widget.NewButton("启动", func() {
		if n.Running {
			return
		}
		n.Mux.Lock()
		defer n.Mux.Unlock()
		if err := n.start(ctx); err != nil {
			d := dialog.NewError(err, util.GetAppWindows(ctx))
			d.Show()
			return
		}
		n.ChangeStatus(util.Running, util.Blue)
	})
}

// start 启动nginx内部封装
func (n *Nginx) start(ctx context.Context) error {
	if err := exec.GetNginxExec().Start(util.GetLocalPath() + `\server\nginx\`); err != nil {
		return err
	}
	n.Running = true
	return nil
}

// Stop 停止nginx
func (n *Nginx) Stop(ctx context.Context) *widget.Button {
	return widget.NewButton("停止", func() {
		if !n.Running {
			return
		}
		n.Mux.Lock()
		defer n.Mux.Unlock()
		if err := n.stop(ctx); err != nil {
			d := dialog.NewError(err, util.GetAppWindows(ctx))
			d.Show()
			return
		}
		n.ChangeStatus(util.Stop, util.Red)
	})
}

// stop 停止nginx内部封装
func (n *Nginx) stop(ctx context.Context) error {
	if err := exec.GetNginxExec().Stop(util.GetLocalPath() + `\server\nginx\`); err != nil {
		return err
	}
	n.Running = false
	return nil
}

// Restart  重启nginx
func (n *Nginx) Restart(ctx context.Context) *widget.Button {
	return widget.NewButton(util.RestartBtn, func() {
		if err := exec.GetNginxExec().Restart(util.GetLocalPath() + `\server\nginx\`); err != nil {
			d := dialog.NewError(err, util.GetAppWindows(ctx))
			d.Show()
			return
		}
		n.ChangeStatus(util.Running, util.Blue)
		n.Running = true
	})
}

// Status 返回当前nginx运行状态
func (n *Nginx) Status() *canvas.Text {
	return n.status
}
