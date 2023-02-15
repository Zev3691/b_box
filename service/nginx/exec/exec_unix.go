//go:build unix

package exec

import (
	"os/exec"
)

type Nginx struct {
}

func GetNginxExec() Nginx {
	return Nginx{}
}

func (n Nginx) Start(baseDir string) error {
	if err := start(n.StartCmd(baseDir)); err != nil {
		return err
	}
	return nil
}

func (n Nginx) Stop(baseDir string) error {
	if err := stop(n.StopCmd(baseDir)); err != nil {
		return err
	}
	return nil
}

func (n Nginx) Restart(baseDir string) error {
	if err := restart(n.RestartCmd(baseDir)); err != nil {
		return err
	}
	return nil
}

func (n Nginx) StartCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh")
	cmd.Dir = baseDir
	return cmd
}

func (n Nginx) StopCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh")
	cmd.Dir = baseDir
	return cmd
}

func (n Nginx) RestartCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("sh")
	cmd.Dir = baseDir
	return cmd
}
