# Demo Flow - Step 1

[prev](../step0/README.md) | [next](../step2/README.md)

## Add the [`pprof`](https://pkg.go.dev/net/http/pprof) endpoints

The demo server is using the default HTTP server mux.
1. Add the following import:
   ```go
   import _ "net/http/pprof"
   ```
2. Restart the server
3. Open the following URL in the browser to see the `pprof` page:  
   `http://localhost:8000/debug/pprof`
