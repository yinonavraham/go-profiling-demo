# Demo Flow - Step 0

[prev](../../README.md) | [next](../step1/README.md)

## Demo preparation and first run

1. Create a dummy file of 1MB under the `data` directory. 
   E.g. use the following to create the directory and a file named `test-1mb`:
   ```
   mkdir data
   mkfile 1m data/test-1mb
   ```
2. Open a terminal and run the demo server:
   ```
   go run .
   ```
   The output should be:
   ```
   Starting server, listening on address: localhost:8000
   ```
3. Open another terminal and run the benchmark tool - 100 threads, 100 connections, for 7 seconds:
   ```
   wrk -t100 -d7s -c100 http://localhost:8000/file/test-1mb 
   ```
   The output should be something like:
   ```
   Running 7s test @ http://localhost:8000/file/test-1mb
     100 threads and 100 connections
     Thread Stats   Avg      Stdev     Max   +/- Stdev
       Latency   157.73ms   68.70ms 562.68ms   90.70%
       Req/Sec     7.07      2.88    20.00     40.15%
     4380 requests in 7.10s, 4.35GB read
     Socket errors: connect 0, read 102, write 0, timeout 0
   Requests/sec:    617.21
   Transfer/sec:    627.80MB
   ```

   ----

   Continue to [step 1](../step1/README.md) to see how to add the pprof endpoints.