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

func (this HostsItem) ToString() string {
	var res string = ""
	if this.ip != "" && this.hosts != "" {
		res = this.ip + " " + this.hosts
	}
	if this.hasComments {
		if this.ip != "" && this.hosts != "" {
			res = res + " " + this.comments + "\n"
		} else {
			res = this.comments
		}
	}
	return res
}
