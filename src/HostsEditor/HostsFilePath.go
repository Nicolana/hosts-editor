package hostseditor

import "runtime"

const DEFAULT_HOSTS_FILE_PATH string = "/etc/hosts"
const WIN_HOSTS_FILE_PATH string = "C:/Windows/System32/drivers/etc/hosts"

func GetHostsFilePath() string {
	sysType := runtime.GOOS

	if sysType == "windows" {
		return WIN_HOSTS_FILE_PATH
	}
	return DEFAULT_HOSTS_FILE_PATH
}
