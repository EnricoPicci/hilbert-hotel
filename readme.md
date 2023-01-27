# Hilbert's Hotel

This repository contains concurrent and non concurrent Go implementations of the Hilbert's Hotel algorithm.

The main purpose is to show how concurrency is the result of design.

![Hilbert's Hotel](./assets/img/hilbert-hotel.png?raw=true).

## Hilbert’s Hotel

[Hilbert’s Hotel](https://youtu.be/SqRY1Bm8EVs?t=1041) is a problem about infinity.

Imagine Hilbert is the owner of an Hotel which has an infinite number of rooms.

One day a bus arrives at the Hilbert’s Hotel. This is a bus with an infinite number of passengers who all want to stay at the Hilbert’s Hotel. Since Hilbert’s Hotel has an infinite number of rooms, there is no problem to accommodate an infinite number of guests.

However, one day something new happens. That day an infinite number of buses, each with an infinite number of passengers, arrive at Hilbert’s Hotel and all passengers of all buses want to stay overnight at Hilbert’s Hotel.

HIlbert thinks a bit and then says “No problem, we can accommodate all of you”. He starts drawing a triangle of numbers and hands over the keys to all passengers following the diagonals of the triangle.

![Hilbert's triangle](./assets/img/hilbert-triangle.png?raw=true).

## A concurrent recursive implementation

To build a concurrent recursive algorithm, let's imagine that Hilbert has some clerks working at the Hotel, actually an infinite number of clerks.

Each clerk will take care of handing over the keys for all the passengers of one bus. The clerk that takes care of the first bus (Bus 1 Clerk) has next the clerk that takes care of the second bus (Bus 2 Clerk) who has next the clerk that takes care of the third bus (Bus 3 Clerk) and so on.

Then there is one more clerk, a clerk who has the task to hand over sequentially all the keys of all the rooms (the Room Key clerk) to the first clerk (Bus 1 Clerk). First room 1, then room 2, then room 3 and so on.

The Bus 1 Clerk (the clerk who receives the keys from the Room Key clerk) knows that the first key he will receive is for the first passenger of its bus. So Bus 1 Clerk will prepare the welcome kit containing the key for the first passenger of the first bus and put it on a desk. Bus 1 Clerk also knows that the second key that it will receive is not for its bus but has to be passed to the next clerk. The third key is for the second passenger of Bus 1, so it is managed by Bus 1 Clerk, but the fourth and fifth keys are not for Bus 1 and so Bus 1 Clerk passes them to the next clerk. When all kits for all the passengers of Bus 1 are ready (we have to imagine that we set a limit to the number of guests, to avoid a never ending program), Bus 1 Clerk will return all of them to Hilbert.

The next clerk, Bus 2 Clerk, behaves in the same way as the first clerk. It knows that the first key it will receive is for the first passenger of its bus (the second bus, Bus 2) while the second key will be for the next clerk, and so on and so forth.

![Clerks working](./assets/img/clerks-working.png?raw=true).

At a certain point, since our program can not go on forever as in Hilbert's case, we will have to stop handing out keys and we will have to signal to all clerks that they have to stop working and that the welcome kits they have prepared have to be sent to Hilbert. Eventually, when the last clerk is told to stop working, it will communicate to Hilbert that also he, Hilbert, can stop working and get a well deserved rest.

![Termination of the work](./assets/img/terminating.png?raw=true).

## Concurrent recursive implementation with Go

[The concurrent implementation](./hilberthotel-concurrent-recursive/) of this algorithm in Go is pretty straightforward.

- [Hilbert](./hilberthotel-concurrent-recursive/hilbert.go) is the main goroutine which launches the entire process and collects the welcome kits created by the bus clerks.
- The [Key Handler Clerk](./hilberthotel-concurrent-recursive/room-key-clerk.go) is a goroutine, launched by Hilbert, which generates sequentially the series of room keys and passes each key to the Bus 1 Clerk until the upper limit of number of keys is reached.
- The [Bus 1 Clerk](./hilberthotel-concurrent-recursive/bus-clerk.go) is another goroutine launched by Hilbert. It receives the keys from the Key Handler Clerk via a channel and implements its logic which is: “prepare the welcome kit for the passengers of your bus and pass the keys that are not for your bus to the next clerk”.
- The [Bus 2 Clerk](./hilberthotel-concurrent-recursive/bus-clerk.go) is another goroutine launched by the Bus 1 clerk. It behaves the same as Bus 1 Clerk.
- All the Bus Clerks behave the same and therefore, while all being executed each as a separate goroutine, they all share the [same code](./hilberthotel-concurrent-recursive/bus-clerk.go).

## Other implementations

Other implementations are provided for comparison purposes.

### Non-concurrent non-recursive implementation

A [non-concurrent non-recursive implementation](./hilberthotel-nonrecursive/) is based on a 2-steps approach.

- first create a slice of int slices that represents the Hilbert triangle
- then loop through the slice of int slices to generate the welcome kits for the various passengers of the various buses

### Non-concurrent recursive implementation

A [non-concurrent recursive implementation](./hilberthotel-closure-recursive/) is based on the same basic design ideas used to build the recursive concurrent solution.

The Room Key clerk is turned into a for loop that generates keys and pass them to the first Bus Clerk.

Each Bus Clerk is implemented as a closure, i.e. a function that holds some state (the same counters count, passengerNumber and nextClerk we have seen in the concurrent code).

Hilbert is the function that triggers the entire execution and collects the welcome kits built by the various Bus Clerks (closures).

### Concurrent implementation with statistics on the use of the buffers of the channels

The channels used to communicate among the various goroutines (Clerks) may be buffered (the buffer size is a parameter that can be passed).

Buffers can speed up the execution of the algorithm since they mitigate the need of synchronization among goroutines (any time 2 goroutines have to be synchronized with a send/receive operation pair on a channel, one goroutine has to wait for the other and therefore is not active and looses CPU cycles).

There is an [implementation of the concurrent algorithm](./hilberthotel-concurrent-recursive-queue-lengths/) that collects statistics about the use of the buffers.

To run this implementation use the following command

`go run ./hilberthotel-concurrent-recursive-queue-lengths/. -numGuests 10000 -buffer 100`

`-numGuests` and `-buffer` are the flags used to specify the corresponding parameters.

## Benchmarks

Some [benchmarks](./benchmarks/) to compare the various implementations are provided.

More [details](./benchmarks/readme.md) on these benchmarks can be found here.
