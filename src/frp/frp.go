package frp

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io"
	"os"
)

type frp struct {
	pid      *os.Process
	checkPid *os.Process // 守护进程
}

var Frp = new(frp)

// Start 启动Frp
func (f *frp) Start() error {
	env := os.Environ()
	file, err := os.OpenFile(LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0666)
	if err != nil {
		return errors.New(fmt.Sprintf("打开frp日志文件错误: %v\n", err.Error()))
	}
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			file,
			file,
		},
	}
	if pid, err := os.StartProcess(getExecPath(), []string{"frpc", "-c", IniConfigPath}, procAttr); err != nil {
		return err
	} else {
		f.pid = pid // 缓存起来
	}
	return nil
}

// GetLog 读取日志文件
func (f *frp) GetLog() ([]byte, error) {
	logFile, err := os.Open(LogFilePath)
	if err != nil {
		return []byte(""), errors.New("无法读取日志文件")
	}
	defer logFile.Close()
	return io.ReadAll(logFile)
}

// Stop 关闭Frp
func (f *frp) Stop() error {
	err := f.pid.Kill()
	if err != nil {
		return err
	}
	return err
}

// IsAlive 查看frp是否还存活状态
func (f *frp) IsAlive() (bool, error) {
	if f.pid == nil {
		return false, errors.New("错误，FRP未启动")
	}
	return process.PidExists(int32(f.pid.Pid))
}

// Restart 重启Frp
func (f *frp) Restart() error {
	if err := f.Stop(); err != nil {
		return err
	}
	return f.Start()
}
