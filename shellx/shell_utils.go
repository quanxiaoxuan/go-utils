package shellx

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// 将shell命令写入脚本
func ShellsWriteToShell(shellFile string, shells []string) {
	file, err := os.OpenFile(shellFile, os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	shellWriter := bufio.NewWriter(file)
	for _, shell := range shells {
		shellWriter.WriteString(shell)
		shellWriter.WriteString("\n")
	}
	err = shellWriter.Flush()
	if err != nil {
		return
	}
	err = file.Close()
	if err != nil {
		return
	}
	log.Info("====== WRITE INTO SHELL FINISH : ", shellFile, " ======")
}

// 执行shell命令
func ExecCommand(execPath, execShell string) (string, error) {
	fmt.Println(execShell)
	var outInfo bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", execShell)
	cmd.Dir = execPath
	// 设置接收
	cmd.Stdout = &outInfo
	// 执行
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return ``, err
	}
	fmt.Println(outInfo.String())
	return outInfo.String(), nil
}
