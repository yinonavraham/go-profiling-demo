# Demo Flow - Step 6

[prev](../step5/README.md) | [next](../step7/README.md)

## Improvement attempt #2

Instead of allocating a new byte array for each request, let's use a `sync.Pool` to reuse already allocated buffers. Important - make sure to return the buffer to the pool once not in use.

1. Create a global pool of byte arrays
   ```go
   var bufs = sync.Pool{
   	New: func() interface{} {
   		return make([]byte, 1024*10)
      	},
   }
   ```
2. Change from:
   ```go
   var b [1024 * 10]byte
   for {
   ```
   To:
   ```go
   b := bufs.Get().([]byte)
   defer func() {
   	bufs.Put(b)
   }()
   ```

Running the benchmark again - the throughput is more or less the same as before - seems there is no real improvement.
Also, both the CPU profile flame graph and the execution trace graph are quite similar to the first improvement attempt - no real improvement here...

----

Continue to [step 7](../step7/README.md) to the next improvement attempt.