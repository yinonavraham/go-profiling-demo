package fileutil

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func OpenFromRequest(res http.ResponseWriter, req *http.Request) (file *os.File, ok bool) {
	if req.Method != http.MethodGet {
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "Method not allowed: %v\n", req.Method)
		return nil, false
	}
	// This is not a security oriented demo, so no path traversal protection and such...
	reqFilePath := strings.TrimPrefix(req.URL.Path, "/file/")
	filePath := filepath.Join("data", reqFilePath)
	var err error
	if file, err = os.Open(filePath); err != nil {
		res.Header().Set("Content-Type", "text/plain")
		if errors.Is(err, os.ErrNotExist) {
			res.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(res, "File not found")
		} else {
			res.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(res, "Error: %v\n", err)
		}
		return nil, false
	}
	return file, true
}
