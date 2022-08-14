package main

import (
	"github.com/Nicolana/hosts-editor/src/frp"
	hostseditor "github.com/Nicolana/hosts-editor/src/hosts-editor"
	"github.com/Nicolana/hosts-editor/src/server"
)

func main() {
	var _ error
	// 初始化Editor
	hostseditor.Editor.FileHandler, _ = hostseditor.Editor.LoadFile()
	err := frp.FrpcIni.LoadFile()
	if err != nil {
		return
	}
	server.RunServer()
}
