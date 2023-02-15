package util

import (
	"b_box/util/log"
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"os"
)

// 状态全局变量
const (
	Stop    = "停止中"
	Running = "运行中"
)

// 按钮全局变量
const (
	StartBtn   = "启动"
	StopBtn    = "停止"
	RestartBtn = "重启"
)

const (
	SkipMysqlAuth = "skip-grant-tables"
)

var localPath string

func SetLocalPath(path string) {
	log.Println("设置本地路径 ", path)
	localPath = path
}

func GetLocalPath() string {
	return localPath
}

// GetAppWindows 从ctx中获取fyne.Windows实例
func GetAppWindows(ctx context.Context) fyne.Window {
	win := ctx.Value("win")
	return win.(fyne.Window)
}

func MyIniFmt(baseDir string, skip string) string {
	return fmt.Sprintf(`[mysqld]
port=3306
basedir=%s
datadir=%s\Data
log-error=%s\Logs.logs
max_connections=200
max_connect_errors=10
character-set-server=utf8mb4
default-storage-engine=INNODB
default_authentication_plugin=mysql_native_password
sql-mode='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION'
%s

[mysql]
default-character-set=utf8mb4
local_infile=ON

[client]
port=3306
default-character-set=utf8mb4
	`, baseDir, baseDir, baseDir, skip)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
