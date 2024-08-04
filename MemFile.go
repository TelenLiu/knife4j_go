package knife4j_go

import (
	"io"
	"os"
	"time"
)

type MemFile struct {
	name    string
	content []byte
	offset  int64
}

type MemFileInfo struct {
	name string
	size int64
}

func (fi MemFileInfo) Name() string       { return fi.name }
func (fi MemFileInfo) Size() int64        { return fi.size }
func (fi MemFileInfo) Mode() os.FileMode  { return 0444 } // Read-only
func (fi MemFileInfo) ModTime() time.Time { return time.Now() }
func (fi MemFileInfo) IsDir() bool        { return false }
func (fi MemFileInfo) Sys() interface{}   { return nil }

func (f *MemFile) Stat() (os.FileInfo, error) {
	return MemFileInfo{name: f.name, size: int64(len(f.content))}, nil
}

func (f *MemFile) Read(b []byte) (n int, err error) {
	if f.offset >= int64(len(f.content)) {
		return 0, io.EOF
	}
	n = copy(b, f.content[f.offset:])
	f.offset += int64(n)
	return n, nil
}

func (f *MemFile) Close() error {
	f.content = nil
	f.offset = 0
	return nil
}

func NewMemFile(name string, content []byte) *MemFile {
	return &MemFile{name: name, content: content}
}
