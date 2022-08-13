package hostseditor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type HostsEditor struct {
	// 文件指针，指向文件
	File *os.File
	// 读取以后的文件对象，可用于修改文件内容
	FileHandler HostsFile
}

// 从文件中加载配置
func (editor HostsEditor) LoadFile() (HostsFile, error) {
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
		var text string = TrimExtraSpace(strings.TrimSpace(str))
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
	return editor.FileHandler, nil
}

// 将配置写回hosts
func (editor HostsEditor) WriteFile() {
	lines := editor.FileHandler.Lines
	// 升序排序
	sort.Slice(lines, func(prev, next int) bool { return lines[prev].index < lines[next].index })
	file, err := os.OpenFile(GetHostsFilePath(), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("写入文件错误 = %v\n", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, value := range lines {
		// fmt.Println(value.index)
		writer.WriteString(value.ToString())
	}
	writer.Flush()
}

func (editor HostsEditor) PrintLines() {
	fmt.Printf("%v\n", editor.FileHandler.Lines)
}

func (editor HostsEditor) PrintByIndex(index int) {
	fmt.Printf("%v\n", editor.FileHandler.Lines[index])
}
