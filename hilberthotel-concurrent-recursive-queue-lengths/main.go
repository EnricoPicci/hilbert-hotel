package main

import (
	"flag"
	"fmt"
)

func main() {
	numGuests := flag.Int("numGuests", 1000, "Number of guests that want to stay at Hilbert's Hotel")
	buffer := flag.Int("buffer", 10, "Size of the buffer to use for the channels")

	flag.Parse()

	fmt.Printf("Run with buffer size=%v and for %v passengers\n\n", *buffer, *numGuests)

	_, queueLengths := Hilbert(*numGuests, *buffer)

	for _, ql := range queueLengths {
		fmt.Printf("Bus %v \t - \t avgLen: %.2f \t - \t stdDevLen: %.2f \t - \t numOfKeysReceived: %v\n", ql.busNumber, ql.avgLen, ql.stdDevLen, len(ql.lengths))
	}
}
