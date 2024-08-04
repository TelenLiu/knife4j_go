package knife4j_go

import (
	"embed"
	"io/fs"
)

//go:embed knife4j-vue-dist/*
var knife4j_dist embed.FS

// 按路径读取文件
func GetKnife4jVueDistFile(filePath string) ([]byte, error) {
	return fs.ReadFile(knife4j_dist, "knife4j-vue-dist/"+filePath)
}

// 返回丢弃外层文件夹的embed.FS
func GetKnife4jVueDistRoot() KnFS {
	f, _ := fs.Sub(knife4j_dist, "knife4j-vue-dist")
	return KnFS{f}
}

// 封装fs.FS，用于读取文件
type KnFS struct {
	fs.FS
}

// 返回文件
func (fs KnFS) Open(name string) (fs.File, error) {
	if content, ok := tempFile[name]; ok {
		return NewMemFile(name, content), nil
	}
	return fs.FS.Open(name)
}

// 保存自定义的fs.File
var tempFile = map[string][]byte{}

// 保存自定义的内容
func SetDiyContent(name string, content []byte) {
	if name != "" && content != nil {
		tempFile[name] = content
	}
}
