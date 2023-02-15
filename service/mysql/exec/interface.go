package exec

import "os/exec"

// Ifc mysql需要实现的具体方法
type Ifc interface {
	Start(baseDir string) error
	Stop(baseDir string) error
	InitCmd(baseDir string) *exec.Cmd
	InitialCmd(baseDir string) *exec.Cmd
	SetPwdCmd(baseDir string) *exec.Cmd
	EnvCmd(baseDir string) *exec.Cmd
	StartCmd(baseDir string) *exec.Cmd
	StopCmd(baseDir string) *exec.Cmd
}
