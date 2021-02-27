package goloader

import "os"

type GoFileNode struct {
	path     string
	FileType string
	File     os.FileInfo
	Parent   *GoFileNode
	Childes  []*GoFileNode
}

func (self *GoFileNode) Path() string {
	return self.path
}
func (self *GoFileNode) Name() string {
	return self.File.Name()
}
