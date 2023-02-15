//go:build unix

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

func (m Mysql) Start(baseDir string) error {
	if err := start(m.StartCmd(baseDir)); err != nil {
		return err
	}
	return nil
}

func (m Mysql) Stop(baseDir string) error {
	if err := stop(m.StopCmd(baseDir)); err != nil {
		return err
	}
	return nil
}

func (m Mysql) InitCmd(baseDir string) *exec.Cmd {
	log.Println("unix")
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
