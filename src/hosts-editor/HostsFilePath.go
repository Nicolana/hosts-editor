package hostseditor

import "runtime"

const DefaultHostsFilePath string = "/etc/hosts"
const WinHostsFilePath string = "C:/Windows/System32/drivers/etc/hosts"

func GetHostsFilePath() string {
	sysType := runtime.GOOS

	if sysType == "windows" {
		return WinHostsFilePath
	}
	return DefaultHostsFilePath
}
