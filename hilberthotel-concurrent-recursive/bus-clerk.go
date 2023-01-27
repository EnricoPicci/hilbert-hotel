package hilberthotelconcurrentrecursive

import (
	"time"

	"github.com/EnricoPicci/hilberthotel"
)

func BusClerk(busNumber int, roomKeysCh <-chan int, welcomeKitsCh chan<- []hilberthotel.WelcomeKit, buffer int, delay time.Duration) {
	var count = 0
	var keyPosition = 0
	var nextClerkCh chan int

	welcomeKits := []hilberthotel.WelcomeKit{}

	for roomKey := range roomKeysCh {
		if nextClerkCh == nil {
			nextClerkCh = make(chan int, buffer)
			go BusClerk(busNumber+1, nextClerkCh, welcomeKitsCh, buffer, delay)
		}
		if count == keyPosition {
			kit := hilberthotel.NewWelcomeKit(busNumber, keyPosition, roomKey, delay)
			welcomeKits = append(welcomeKits, kit)
			keyPosition++
			count = 0
			continue
		}
		nextClerkCh <- roomKey
		count++
	}

	if nextClerkCh != nil {
		welcomeKitsCh <- welcomeKits
		close(nextClerkCh)
	} else {
		close(welcomeKitsCh)
	}
}
