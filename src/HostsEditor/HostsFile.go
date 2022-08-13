package hostseditor

type HostsItem struct {
	index        int // 标记行数
	hosts        string
	ip           string
	hasComments  bool   // 标记当前行有没有注释
	comments     string // 存放注释
	originString string // 存放原始字符串
}

type HostsFile struct {
	Lines []HostsItem
}

func (item HostsItem) ToString() string {
	var res string = ""
	if item.ip != "" && item.hosts != "" {
		res = item.ip + " " + item.hosts
	}
	if item.hasComments {
		if item.ip != "" && item.hosts != "" {
			res += " " + item.comments + GetReturnSymbol()
		} else {
			res = item.comments
		}
	} else {
		res += GetReturnSymbol()
	}
	return res
}
