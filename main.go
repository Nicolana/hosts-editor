package main

import (
	hostseditor "github.com/Nicolana/hosts-editor/src/hosts-editor"
	server "github.com/Nicolana/hosts-editor/src/server"
)

func main() {
	var _ error
	// 初始化Editor
	hostseditor.Editor.FileHandler, _ = hostseditor.Editor.LoadFile()
	server.RunServer()
}
