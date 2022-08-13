package hostseditor

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/Nicolana/hosts-editor/src/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)

type HostsEditor struct {
	// 文件指针，指向文件
	File *os.File
	max  int
	// 读取以后的文件对象，可用于修改文件内容
	FileHandler HostsFile
}

// LoadFile 从文件中加载配置
func (editor *HostsEditor) LoadFile() (HostsFile, error) {
	var err error
	editor.File, err = os.Open(GetHostsFilePath())
	editor.FileHandler = HostsFile{}
	editor.FileHandler.Lines = []HostsItem{}

	if err != nil {
		fmt.Println("Hosts文件打开失败!")
		return editor.FileHandler, err
	}
	defer editor.File.Close()
	reader := bufio.NewReader(editor.File)
	index := 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		index++

		// 切分出ip、host、comment, 如果带#符号，说明有注释
		var text = TrimExtraSpace(strings.TrimSpace(str))
		if text == "" {
			continue
		}
		var comment string
		var host string
		var ip string
		i := strings.Index(str, "#")
		if i > -1 {
			// 移除多余的空白字符串
			text = TrimExtraSpace(strings.TrimSpace(str[:i]))
			comment = str[i:]
		}
		s := strings.Split(text, " ")
		if len(s) > 1 {
			ip = s[0]
			host = s[1]
		}
		// 如果是Hosts内容，则记录在Lines里面
		row := &HostsItem{
			index:        index,
			hosts:        host,
			ip:           ip,
			hasComments:  i > -1,
			comments:     comment,
			originString: str,
		}
		editor.FileHandler.Lines = append(editor.FileHandler.Lines, *row)
	}
	editor.max = index
	return editor.FileHandler, nil
}

// WriteFile 将配置写回hosts
func (editor *HostsEditor) WriteFile() error {
	lines := editor.FileHandler.Lines
	// 升序排序
	//sort.Slice(lines, func(prev, next int) bool { return lines[prev].index < lines[next].index })
	file, err := os.OpenFile(GetHostsFilePath(), os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0666)
	if err != nil {
		return errors.New(fmt.Sprintf("写入文件错误: %v\n", err.Error()))
	}
	defer file.Close()
	for _, value := range lines {
		var line = value.ToString()
		if _, err := file.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

func (editor *HostsEditor) PrintLines() {
	fmt.Printf("%v\n", editor.FileHandler.Lines)
}

func (editor *HostsEditor) PrintByIndex(index int) {
	fmt.Printf("%v\n", editor.FileHandler.Lines[index])
}

// GetLines 分页获取hosts数据
func (editor *HostsEditor) GetLines(page, size int, search string) gin.H {
	//var total = len(editor.FileHandler.Lines)
	lines := utils.Filter(editor.FileHandler.Lines, func(item HostsItem) bool {
		if strings.Contains(item.hosts, search) {
			return true
		}
		return false
	})
	var list []gin.H
	for _, value := range lines {
		if !value.hasComments || (value.ip != "" && value.hosts != "") {
			list = append(list, value.ToMap())
		}
	}
	// 处理分页功能
	var total = len(list)
	if total > 1 {
		var left = (page - 1) * size
		var right = page * size
		if right >= len(list) {
			right = len(list)
		}
		fmt.Printf("left = %d, right = %d \n", left, right)
		list = list[left:right]
	}

	return gin.H{
		"page":  page,
		"size":  size,
		"total": total,
		"data":  list,
	}
}

// UpdateLineByIndex 通过Index更新Hosts数据
func (editor *HostsEditor) UpdateLineByIndex(index int, hosts, ip, comments string) (HostsItem, error) {
	var h HostsItem
	if index > len(editor.FileHandler.Lines)-1 {
		return h, errors.New("数据越界了")
	}
	var line = &editor.FileHandler.Lines[index]
	if line.hasComments && line.ip == "" && line.hosts == "" {
		return h, errors.New("该行为注释行，无法修改")
	}
	line.hosts = hosts
	line.ip = ip
	line.comments = comments
	line.hasComments = comments != ""
	//fmt.Println("line =", line)
	err := editor.WriteFile()
	if err != nil {
		return HostsItem{}, err
	}
	return *line, nil
}

// DeleteLineByIndex 删除hosts
func (editor *HostsEditor) DeleteLineByIndex(index int) (HostsItem, error) {
	var item HostsItem
	if index > len(editor.FileHandler.Lines)-1 {
		return item, errors.New("数据越界了")
	}
	fmt.Println("原始长度", len(editor.FileHandler.Lines))
	item = editor.FileHandler.Lines[index]
	lines := editor.FileHandler.Lines
	lines = append(lines[:index], lines[index+1:]...)
	(&editor.FileHandler).Lines = lines

	fmt.Println("删除后长度", len(editor.FileHandler.Lines))
	err := editor.WriteFile()
	if err != nil {
		return HostsItem{}, err
	}
	return item, nil
}

// HostsExists 检查hosts是否存在
func (editor *HostsEditor) HostsExists(hosts string) bool {
	for _, value := range editor.FileHandler.Lines {
		if value.hosts == hosts {
			return true
		}
	}
	return false
}

// IpExists 检查IP是否存在
func (editor *HostsEditor) IpExists(hosts string) bool {
	for _, value := range editor.FileHandler.Lines {
		if value.hosts == hosts {
			return true
		}
	}
	return false
}

// AddNewRow 新增Hosts解析
func (editor *HostsEditor) AddNewRow(hosts, ip, comments string) (HostsItem, error) {
	if editor.HostsExists(hosts) {
		return HostsItem{}, errors.New("该域名已存在")
	}
	var item = &HostsItem{
		index:       editor.max + 1,
		hosts:       hosts,
		ip:          ip,
		hasComments: comments != "",
		comments:    comments,
	}
	editor.FileHandler.Lines = append(editor.FileHandler.Lines, *item)
	err := editor.WriteFile()
	if err != nil {
		return HostsItem{}, err
	}
	return *item, nil
}

var Editor = new(HostsEditor)
