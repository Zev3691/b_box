//go:build windows

package exec

import (
	"fmt"
	"log"
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
	log.Println("windows")
	cmd := exec.Command("powershell", `.\mysqld.exe`, "--install", "mysql")
	cmd.Dir = baseDir + `\bin`
	return cmd
}

func (m Mysql) InitialCmd(baseDir string) *exec.Cmd {
	cmd := exec.Command("powershell", `.\mysqld.exe`, "--initialize", "--console")
	cmd.Dir = baseDir + `\bin`
	return cmd
}

func (m Mysql) SetPwdCmd(baseDir string) *exec.Cmd {
	sqlPath := baseDir + "\\init.sql"
	cmd := exec.Command("powershell", fmt.Sprintf(`.\mysql -u root -p -e "source %s"`, sqlPath))
	cmd.Dir = baseDir + `\bin` // mysql执行文件路径
	return cmd
}

func (m Mysql) EnvCmd(baseDir string) *exec.Cmd {
	c := `setx PATH "%PATH%;` + fmt.Sprintf(`%s\bin"`, baseDir)
	log.Println("set path ", c)
	return exec.Command("powershell", c)
}

func (m Mysql) StartCmd(baseDir string) *exec.Cmd {
	return exec.Command("powershell", "net", "start", "mysql")
}

func (m Mysql) StopCmd(baseDir string) *exec.Cmd {
	return exec.Command("powershell", "net", "stop", "mysql")
}
