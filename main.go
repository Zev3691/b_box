package main

import (
	"b_box/service"
	"os"
)

func init() {
	// 设置主题为明亮
	_ = os.Setenv("FYNE_THEME", "light")
	// 设置字体,解决中文乱码问题
	_ = os.Setenv("FYNE_FONT", "util/font/AlibabaPuHuiTi-2-45-Light.ttf")
}

func main() {
	service.Do()
}
