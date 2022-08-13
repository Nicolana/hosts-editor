package hostseditor

import (
	"regexp"
	"runtime"
)

func TrimExtraSpace(str string) string {
	// 移除多余空白
	reg, _ := regexp.Compile(`\s+`)
	// 如果有两个极其以上的空白，直接替换成单个空白
	return reg.ReplaceAllString(str, " ")
}

func GetReturnSymbol() string {
	sysType := runtime.GOOS
	if sysType == "windows" {
		return "\r\n"
	}
	return "\n"
}
