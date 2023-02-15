package menu

import (
	"b_box/service/mysql/exec"
	"b_box/util"
	"b_box/util/log"
	"context"
	"errors"
	"fyne.io/fyne/v2"
	"io/fs"

	"path/filepath"
	"sync"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Menu只运行调用其他包的方法，不允许其他包调用menu

type Menu struct {
	Mux *sync.RWMutex
}

var once sync.Once
var menuInstance *Menu

func init() {
	newMenu()
}

func newMenu() *Menu {
	return &Menu{
		Mux: &sync.RWMutex{},
	}
}

func getMenuInstance() *Menu {
	if menuInstance == nil {
		once.Do(
			func() {
				menuInstance = newMenu()
			},
		)
	}

	return menuInstance
}

func GetMenuCtl() *Menu {
	return getMenuInstance()
}

func (m *Menu) Setting(ctx context.Context) *widget.Button {
	window := util.GetAppWindows(ctx)
	return widget.NewButton("设置", func() {
		d := dialog.NewCustom("设置", "关闭", m.setting(ctx), window)
		d.Show()
	},
	)
}

func (m *Menu) setting(ctx context.Context) *fyne.Container {
	window := util.GetAppWindows(ctx)
	return container.NewGridWithRows(1,
		widget.NewButton("初始化数据库",
			func() {
				path := filepath.ToSlash(util.GetLocalPath() + `\server\mysql\`)
				err := exec.Init(path, exec.Mysql{})
				if err != nil && !errors.Is(err, fs.ErrExist) {
					log.Println(err)
					dialog.ShowError(err, window)
				} else {
					dialog.NewCustom(
						"",
						"关闭",
						canvas.NewText("初始化成功", util.Blue),
						window).Show()
				}
			},
		),
	)
}
