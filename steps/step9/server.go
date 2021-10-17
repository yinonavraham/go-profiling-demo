package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/yinonavraham/go-profiling-demo/fileutil"

	"expvar" // Import the expvar package
	_ "net/http/pprof"
)

// Define the exposed variables
var concurrentRequests expvar.Int
var totalRequests expvar.Int
var totalBytes expvar.Int

func main() {
	// Publish a map variable with the metrics values
	metrics := expvar.NewMap("metrics")
	metrics.Set("concurrentRequests", &concurrentRequests)
	metrics.Set("totalRequests", &totalRequests)
	metrics.Set("totalBytes", &totalBytes)

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
	// Add to request metrics counters
	totalRequests.Add(1)
	concurrentRequests.Add(1)
	defer concurrentRequests.Add(-1)

	file, ok := fileutil.OpenFileHandleFromRequest(res, req)
	if !ok {
		return
	}
	defer file.Close()
	res.Header().Set("Content-Type", "application/octet-stream")
	res.WriteHeader(http.StatusOK)
	b := bufs.Get().([]byte)
	defer func() {
		bufs.Put(b)
	}()
	for {
		n, err := file.Read(b[:])
		if err != nil {
			break
		}
		// Count the bytes read
		totalBytes.Add(int64(n))
		res.Write(b[:n])
	}
}
