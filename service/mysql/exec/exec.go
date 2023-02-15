package exec

import (
	"b_box/util"
	"b_box/util/log"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// writeInI 写配置文件
func writeInI(baseDir string, skip bool) error {
	if baseDir == "" {
		return errors.New("必须在配置文件指定base目录!")
	}
	fd, err := os.OpenFile(baseDir+"\\my.ini", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("创建ini文件失败: %v", err)
	}
	if skip {
		_, err = fd.WriteString(util.MyIniFmt(baseDir, util.SkipMysqlAuth))
		if err != nil {
			return fmt.Errorf("写入ini文件失败: %v", err)
		}
	} else {
		_, err = fd.WriteString(util.MyIniFmt(baseDir, ""))
		if err != nil {
			return fmt.Errorf("写入ini文件失败: %v", err)
		}
	}

	return nil
}

// initMysql 等于执行 ./mysqld --install
func initMysql(cmd *exec.Cmd) error {

	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出 标准错误
	err := cmd.Run()

	log.Println("安装mysql服务 cmd.String() %v", cmd.String())
	log.Println("安装mysql服务 cmd.dir %v", cmd.Dir)
	log.Println("安装mysql服务 out %v", stdout.String())
	log.Println("安装mysql服务 err %v", stderr.String())

	if err != nil {
		log.Println("安装mysql服务失败: %v", err)
		return fmt.Errorf("安装mysql服务失败: %v", err)
	}
	return nil
}

// initialMysql 等于执行 ./mysqld --initialize --console
func initialMysql(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出 标准错误
	err := cmd.Run()

	log.Println("初始化mysql数据库 cmd.String() ", cmd.String())
	log.Println("初始化mysql数据库 out ", stdout.String())
	log.Println("初始化mysql数据库 err ", stderr.String())

	if err != nil {
		log.Println("初始化mysql数据库失败: ", err)
		return fmt.Errorf("初始化mysql数据库失败: %v", err)
	}
	return nil
}

// setENV 设置mysql环境变量
func setENV(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出 标准错误
	err := cmd.Run()

	log.Println("增加环境变量 cmd.String() ", cmd.String())
	log.Println("增加环境变量 out ", stdout.String())
	log.Println("增加环境变量 err ", stderr.String())

	if err != nil {
		log.Println("增加环境变量: %v", err)
		return fmt.Errorf("增加环境变量: %v", err)
	}
	return nil
}

// resetPwd 重置mysql初始密码
func resetPwd(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出 标准错误
	err := cmd.Run()

	log.Println("重载mysql密码 cmd.String() ", cmd.String())
	log.Println("重载mysql密码 out ", stdout.String())
	log.Println("重载mysql密码 err ", stderr.String())

	if err != nil {
		log.Println("重载mysql密码失败: ", err)
		return fmt.Errorf("重载mysql密码失败: %v", err)
	}
	return nil
}

// start 启动mysql
func start(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出, 标准错误
	err := cmd.Run()

	log.Println("启动mysql服务 cmd.String() ", cmd.String())
	log.Println("启动mysql服务 out ", stdout.String())
	log.Println("启动mysql服务 err ", stderr.String())

	if err != nil {
		log.Println("启动mysql服务: %v", err)
		return fmt.Errorf("启动mysql服务失败: %v", err)
	}
	return nil
}

// stop 停止mysql
func stop(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出, 标准错误
	err := cmd.Run()

	log.Println("停止mysql服务 cmd.String() ", cmd.String())
	log.Println("停止mysql服务 out ", string(stdout.Bytes()))
	log.Println("停止mysql服务 err ", string(stderr.Bytes()))

	if err != nil {
		log.Println("停止mysql服务失败: %v", err)
		return fmt.Errorf("停止mysql服务失败: %v", err)
	}

	return nil
}

// Init 给设置按钮内部的初始化数据库按钮调用的封装方法集 见../menu
func Init(baseDir string, ifc Ifc) error {
	exist, err := util.PathExists(baseDir + `\.init`)
	if err != nil {
		log.Println("检查数据库init文件是否存在err %v", err)
	}
	log.Println("检查数据库init文件是否存在 %v", exist)
	if exist {
		return nil
	}

	if err := writeInI(baseDir, true); err != nil {
		return err
	}

	if err := initMysql(ifc.InitCmd(baseDir)); err != nil {
		return err
	}

	if err := initialMysql(ifc.InitialCmd(baseDir)); err != nil {
		return err
	}

	if err := setENV(ifc.EnvCmd(baseDir)); err != nil {
		return err
	}

	if err := start(ifc.StartCmd("")); err != nil {
		return err
	}

	if err := resetPwd(ifc.SetPwdCmd(baseDir)); err != nil {
		return err
	}

	if err := writeInI(baseDir, false); err != nil {
		return err
	}

	if err := stop(ifc.StopCmd("")); err != nil {
		return err
	}

	fd, _ := os.Create(baseDir + `\.init`)
	defer fd.Close()

	return nil
}
