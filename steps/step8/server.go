package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/yinonavraham/go-profiling-demo/fileutil"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/file/", handleGetFile)

	address := "localhost:8000"
	fmt.Printf("Starting server, listening on address: %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

var bufs = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024*10)
	},
}

func handleGetFile(res http.ResponseWriter, req *http.Request) {
	file, ok := fileutil.OpenFileHandleFromRequest(res, req) // Use the FileHandle to collect profiling data
	if !ok {
		return
	}
	defer file.Close()
	res.Header().Set("Content-Type", "application/octet-stream")
	res.WriteHeader(http.StatusOK)
	b := bufs.Get().([]byte)
	for {
		n, err := file.Read(b[:])
		if err != nil {
			break
		}
		res.Write(b[:n])
	}
	bufs.Put(b)
}
