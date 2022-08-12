package hostseditor

import "regexp"

func TrimExtraSpace(str string) string {
	// 移除多余空白
	reg := regexp.MustCompile("\\s+")
	// 如果有两个极其以上的空白，直接替换成单个空白
	return reg.ReplaceAllString(str, " ")
}
