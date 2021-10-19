package fileutil

import (
	"io"
	"net/http"
	"os"
	"runtime/pprof"
)

var profile *pprof.Profile

func init() {
	const profile_name = "profiling_demo_file_handles"
	profile = pprof.Lookup(profile_name)
	if profile == nil {
		// profile = pprof.NewProfile(profile_name)
	}
}

type FileHandle struct {
	file *os.File
}

var _ io.Reader = &FileHandle{}
var _ io.Closer = &FileHandle{}

func (h *FileHandle) Read(b []byte) (int, error) {
	return h.file.Read(b)
}

func (h *FileHandle) Close() error {
	profile.Remove(h)
	return h.file.Close()
}

func OpenFileHandleFromRequest(res http.ResponseWriter, req *http.Request) (fileHandle *FileHandle, ok bool) {
	if file, ok := OpenFileFromRequest(res, req); ok {
		fileHandle = &FileHandle{
			file: file,
		}
		profile.Add(fileHandle, 1)
		return fileHandle, true
	}
	return nil, false
}
