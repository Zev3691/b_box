//go:build linux

package exec

import (
	"b_box/util/log"
	"os/exec"
)

type Mysql struct {
}

func GetMysqlExec() Mysql {
	return Mysql{}
}

func (m Mysql) Start() error {
	return nil
}

func (m Mysql) Stop() error {
	return nil
}

func (m Mysql) InitCmd(baseDir string) *exec.Cmd {
	log.Println("linux")
	cmd := exec.Command("sh", "")
	return cmd
}

func (m Mysql) InitialCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh", "")
	return cmd
}

func (m Mysql) SetPwdCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh", "")
	return cmd
}

func (m Mysql) EnvCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh", "")
	return cmd
}

func (m Mysql) StartCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh", "")
	return cmd
}

func (m Mysql) StopCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh", "")
	return cmd
}
