# Go Profiling Demo
Go application for demoing profiling in Go using standard library tools and packages such as [`pprof`](https://github.com/google/pprof/blob/master/doc/README.md), [`trace`](https://golang.org/cmd/trace/) and [`expvar`](https://pkg.go.dev/expvar).

This is a simple HTTP server application, serving files from a `data` directory. 
It mainly implements the following endpoint:
```
GET /file/<file-path>
```

## Setup
1. Clone this repository
2. Install [wrk](https://github.com/wg/wrk) benchmarking tool  
   (used to send concurrent requests to the demo server application)

## Demo Flow

Following are the steps of this demo.
Follow the links to each step to read what is done in the step and to see the code changes.

0. [Step 0](steps/step0/README.md) - Demo preparation and first run
1. [Step 1](steps/step1/README.md) - Add the [`pprof`](https://pkg.go.dev/net/http/pprof) endpoints
2. [Step 2](steps/step2/README.md) - Collect CPU profile and visualize it
3. [Step 3](steps/step3/README.md) - Use execution tracer to visualize Go routines scheduling, GC events, and more
4. [Step 4](steps/step4/README.md) - Collect memory profile and visualize it
5. [Step 5](steps/step5/README.md) - Improvement attempt #1
6. [Step 6](steps/step6/README.md) - Improvement attempt #2
7. [Step 7](steps/step7/README.md) - Improvement attempt #3
8. [Step 8](steps/step8/README.md) - Add a custom profile
9. [Step 9](steps/step9/README.md) - Expose operational values

----

## Notes

### Don't use the default serve mux

Using the default serve mux complicates the ability to put access control and enables other packages to expose endpoints implicitly, similar to how the `pprof` endpoints are added just by adding an anonymous import.
It is advised to use your own serve mux and add the pprof endpoints explicitly, preferably with an auth middleware. 
For example:
```go
func main() {
   mux := http.NewServeMux()
   mux.HandleFunc("/file/", handleGetFile)
   addPprofHandlers(mux)
   // ...
	log.Fatal(http.ListenAndServe(address, authMiddleware(mux))) 
}

func addPprofHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func authMiddleware(next http.Handler) http.Handler {
   // ...
}
```

----

## Links
* [Go Diagnostics](https://golang.org/doc/diagnostics)
* [pprof tool by Google](https://github.com/google/pprof/blob/master/doc/README.md)