package exec

import (
	"b_box/util/log"
	"bytes"
	"fmt"
	"os/exec"
)

// 程序具体执行的命令集中在这三个函数中

func start(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出, 标准错误
	err := cmd.Run()

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	log.Println("启动nginx服务 cmd.String() ", cmd.String())
	log.Println("启动nginx服务 out ", outStr)
	log.Println("启动nginx服务 err ", errStr)

	if err != nil {
		log.Println("启动nginx服务: %v", err)
		return fmt.Errorf("启动nginx服务: %v", err)
	}
	return nil
}

func stop(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出, 标准错误
	err := cmd.Run()

	log.Println("停止nginx服务 cmd.String() ", cmd.String())
	log.Println("停止nginx服务 out ", string(stdout.Bytes()))
	log.Println("停止nginx服务 err ", string(stderr.Bytes()))

	if err != nil {
		log.Println("停止nginx服务失败: %v", err)
		return fmt.Errorf("停止nginx服务失败: %v", err)
	}

	return nil
}

func restart(cmd *exec.Cmd) error {
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr // 标准输出, 标准错误
	err := cmd.Run()

	log.Println("重启nginx服务 cmd.String() ", cmd.String())
	log.Println("重启nginx服务 out ", string(stdout.Bytes()))
	log.Println("重启nginx服务 err ", string(stderr.Bytes()))

	if err != nil {
		log.Println("重启nginx服务失败: %v", err)
		return fmt.Errorf("重启止nginx服务失败: %v", err)
	}

	return nil
}
