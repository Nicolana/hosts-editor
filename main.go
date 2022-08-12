package main

import (
	hostseditor "hostswitcher/src/HostsEditor"
)

func main() {
	editor := new(hostseditor.HostsEditor)
	var err error
	editor.FileHandler, err = editor.LoadFile()
	if err != nil {
		return
	}
	// editor.PrintByIndex(6)
	editor.PrintByIndex(13)
	// editor.WriteFile()
	// server.RunServer()
}
