package hilberthotelconcurrentrecursive

import (
	"fmt"
	"time"

	"github.com/EnricoPicci/hilberthotel"
)

func Hilbert(upTo int, buffer int, delay time.Duration, verbose bool) []hilberthotel.WelcomeKit {
	if buffer < 0 {
		buffer = 0
	}
	keysCh := make(chan int, buffer)
	go RoomKeysClerk(upTo, keysCh)

	hilbertCh := make(chan []hilberthotel.WelcomeKit, buffer)
	go BusClerk(1, keysCh, hilbertCh, buffer, delay)

	kits := []hilberthotel.WelcomeKit{}
	for busKits := range hilbertCh {
		kits = append(kits, busKits...)
	}

	if verbose {
		fmt.Println()
		fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", len(kits))
	}

	return kits
}
