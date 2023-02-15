package exec

import "os/exec"

// Ifc 定义nginx需要实现的方法,针对不同平台进行实现
type Ifc interface {
	Start(baseDir string) error
	Stop(baseDir string) error
	Restart(baseDir string) error
	StartCmd(baseDir string) *exec.Cmd
	StopCmd(baseDir string) *exec.Cmd
	RestartCmd(baseDir string) *exec.Cmd
}
