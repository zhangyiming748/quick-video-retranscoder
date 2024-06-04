package util

import (
	"path"
	"runtime"
)

var root string = "E:\\pikpak"

func SetRoot() {
	_, filename, _, _ := runtime.Caller(0)
	root = path.Dir(filename)
}

func GetRoot() string {
	return root
}
