# Demo Flow - Step 1

[prev](../step0/README.md) | [next](../step2/README.md)

## Add the [`pprof`](https://pkg.go.dev/net/http/pprof) endpoints

The demo server is using the default HTTP server mux.
1. Add the following import:
   ```go
   import _ "net/http/pprof"
   ```
2. Kill the server (`Ctrl-C`) and start it again (`go run .`)
3. Open the following URL in the browser to see the `pprof` page:  
   `http://localhost:8000/debug/pprof`


(see modified source code: [server.go](server.go))

See the [`net/http/pprof`](https://pkg.go.dev/net/http/pprof) package documentation for further details on the information the `pprof` endpoints provide.

----

Continue to [step 2](../step2/README.md) to see how to use pprof to collect CPU profile and analyse it.