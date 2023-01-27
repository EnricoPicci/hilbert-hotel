# Hiilbert's Hotel benchmarks

There are 3 different implementations of the Hilbert's Hotel algorithm

- non recursive
- recursive using closures
- concurrently recursive using goroutines and channels
  These different implementations are benchmarked in this package.

## Parameters

It is possible to pass some parameters to the benchmarks

- **numGuests**: Number of guests that want to stay at Hilbert's Hotel
- **delayMicrosec**: Delay in microsecs to make a welcome kit (simulates work to be done for each guest)

In addition to these paramenters we can use all the parameters that the `go test` command accepts, e.g. `-benchtime 3s` to set a 3 seconds duration for each benchmark.

## Example

An execution with 10.000 guests and a delay of 1 microsecond where only benchmarks are run (i.e. no tests are run) and each benchmark is run for 2 seconds

`go test -benchmem -bench=. ./benchmarks/. -run none -benchtime 2s -numGuests 10000 -delayMicrosec 1`
