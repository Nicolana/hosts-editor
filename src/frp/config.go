package frp

import (
	"github.com/Nicolana/hosts-editor/src/utils"
	"runtime"
)

const goos = runtime.GOOS

var rootPath, _ = utils.GetRootPath()

func getExecPath() string {
	if goos == "windows" {
		return rootPath + "/lib/bin/frp/windows/frpc.exe"
	}
	if goos == "linux" {
		return rootPath + "/lib/bin/frp/linux/frpc"
	}
	return rootPath + "/lib/bin/frp/darwin/frpc"
}

var IniConfigPath = rootPath + "/lib/bin/frp/conf/frpc.ini"

var LogRootPath = rootPath + "/log"

var LogFilePath = LogRootPath + "/frp-log.txt"
