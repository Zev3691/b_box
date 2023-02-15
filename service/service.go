package service

import (
	"b_box/service/menu"
	"b_box/service/mysql"
	"b_box/service/nginx"
	"b_box/util"
	"context"
	fyne "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"os"
)

func Do() {
	// 获取当前执行目录,所以与项目有关的文件都需要以此文件夹为基础目录
	dir, _ := os.Getwd()
	// 设置为全局
	util.SetLocalPath(dir)

	myApp := fyne.New()
	// 设置标题
	myWin := myApp.NewWindow("B-Box")
	// 设置ctx, 目的是为了后续可以在此窗口的基础上打开新的辅助窗口
	// 也可为后续通过ctx进行log追踪
	ctx := context.WithValue(context.Background(), "win", myWin)
	// 主布局
	content := container.NewGridWithRows(3,
		mysql.GetMysqlCtl().Layout(ctx),
		nginx.GetNginxCtl().Layout(ctx),
		menu.GetMenuCtl().Layout(ctx),
	)
	myWin.SetContent(content)

	myWin.ShowAndRun()
}
