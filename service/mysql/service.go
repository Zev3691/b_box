package mysql

import (
	"context"
	"image/color"
	"sync"

	"b_box/service/mysql/exec"
	"b_box/util"

	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Mysql struct {
	Mux     *sync.RWMutex
	Running bool

	status *canvas.Text
}

var once sync.Once

// 全局mysql操作单例
var mysqlInstance *Mysql

func init() {
	newMysql()
}

func newMysql() *Mysql {
	return &Mysql{
		Mux:     &sync.RWMutex{},
		Running: false,
		status:  canvas.NewText(util.Stop, util.Red),
	}
}

func getMysqlInstance() *Mysql {
	if mysqlInstance == nil {
		once.Do(
			func() {
				mysqlInstance = newMysql()
			},
		)
	}

	return mysqlInstance
}

// GetMysqlCtl 获取全局mysql操作单例
func GetMysqlCtl() *Mysql {
	return getMysqlInstance()
}

// ChangeStatus 改变mysql状态显示
func (m *Mysql) ChangeStatus(nextStatus string, nextStatusColor color.Color) {
	m.status.Text = nextStatus
	m.status.Color = nextStatusColor
	m.status.Refresh()
}

func (m *Mysql) Start(ctx context.Context) *widget.Button {
	return widget.NewButton(util.StartBtn, func() {
		if m.Running {
			return
		}

		m.Mux.Lock()
		defer m.Mux.Unlock()
		if err := m.start(ctx); err != nil {
			d := dialog.NewError(err, util.GetAppWindows(ctx))
			d.Show()
			return
		}

		m.ChangeStatus(util.Running, util.Blue)
	})
}

func (m *Mysql) start(ctx context.Context) error {
	if err := exec.GetMysqlExec().Start(util.GetLocalPath() + `\server\mysql\`); err != nil {
		return err
	}
	m.Running = true
	return nil
}

func (m *Mysql) Stop(ctx context.Context) *widget.Button {
	return widget.NewButton(util.StopBtn, func() {
		if !m.Running {
			return
		}

		m.Mux.Lock()
		defer m.Mux.Unlock()
		if err := m.stop(ctx); err != nil {
			d := dialog.NewError(err, util.GetAppWindows(ctx))
			d.Show()
			return
		}

		m.ChangeStatus(util.Stop, util.Red)
	})
}

func (m *Mysql) stop(ctx context.Context) error {
	if err := exec.GetMysqlExec().Stop(util.GetLocalPath() + `\server\mysql\`); err != nil {
		return err
	}
	m.Running = false
	return nil
}

func (m *Mysql) Status() *canvas.Text {
	return m.status
}

func (m *Mysql) Healthy() {
	// todo
}
