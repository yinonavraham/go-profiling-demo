package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yinonavraham/go-profiling-demo/fileutil"

	_ "net/http/pprof" // New import
)

func main() {
	http.HandleFunc("/file/", handleGetFile)

	address := "localhost:8000"
	fmt.Printf("Starting server, listening on address: %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handleGetFile(res http.ResponseWriter, req *http.Request) {
	file, ok := fileutil.OpenFromRequest(res, req)
	if !ok {
		return
	}
	defer file.Close()
	res.Header().Set("Content-Type", "application/octet-stream")
	res.WriteHeader(http.StatusOK)
	for {
		var b [1024]byte
		n, err := file.Read(b[:])
		if err != nil {
			break
		}
		res.Write(b[:n])
	}
}
