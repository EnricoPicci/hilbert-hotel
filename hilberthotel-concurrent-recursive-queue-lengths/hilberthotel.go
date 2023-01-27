package main

import (
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/EnricoPicci/hilberthotel"
)

func RoomKeysClerk(upTo int, keysCh chan<- int) {
	for i := 0; i < upTo; i++ {
		keysCh <- i + 1
	}
	close(keysCh)
}

type QueueLengths struct {
	lengths   []int
	avgLen    float64
	stdDevLen float64
	busNumber int
}

func BusClerk(busNumber int, roomKeysCh <-chan int, welcomeKitsCh chan<- []hilberthotel.WelcomeKit, queueLengthsCh chan<- QueueLengths, buffer int) {
	delay := 10 * time.Microsecond
	var count = 0
	var passengerNumber = 1
	var nextClerkCh chan int

	queueLengths := QueueLengths{lengths: []int{}, busNumber: busNumber}

	welcomeKits := []hilberthotel.WelcomeKit{}

	for roomKey := range roomKeysCh {
		queueLengths.lengths = append(queueLengths.lengths, len(roomKeysCh))
		count++
		if nextClerkCh == nil {
			nextClerkCh = make(chan int, buffer)
			go BusClerk(busNumber+1, nextClerkCh, welcomeKitsCh, queueLengthsCh, buffer)
		}
		if count == passengerNumber {
			kit := hilberthotel.NewWelcomeKit(busNumber, passengerNumber, roomKey, delay)
			welcomeKits = append(welcomeKits, kit)
			passengerNumber++
			count = 0
			continue
		}
		nextClerkCh <- roomKey
	}

	if nextClerkCh != nil {
		avgLen, stdDevLen := meanStdDev(queueLengths.lengths)
		queueLengths.avgLen = avgLen
		queueLengths.stdDevLen = stdDevLen

		welcomeKitsCh <- welcomeKits
		queueLengthsCh <- queueLengths
		close(nextClerkCh)
	} else {
		close(welcomeKitsCh)
		close(queueLengthsCh)
	}
}

func meanStdDev(intSlice []int) (mean float64, stdDev float64) {
	for _, v := range intSlice {
		mean += float64(v)
	}
	mean = float64(mean) / float64(len(intSlice))

	for _, v := range intSlice {
		stdDev += math.Pow(float64(v)-mean, 2)
	}
	stdDev = math.Sqrt(stdDev / float64(len(intSlice)))

	return mean, stdDev
}

func Hilbert(upTo int, buffer int) ([]hilberthotel.WelcomeKit, []QueueLengths) {
	if buffer < 0 {
		buffer = 0
	}
	keysCh := make(chan int, buffer)
	go RoomKeysClerk(upTo, keysCh)

	hilbertCh := make(chan []hilberthotel.WelcomeKit, buffer)
	queueLengthsCh := make(chan QueueLengths, buffer)
	go BusClerk(1, keysCh, hilbertCh, queueLengthsCh, buffer)

	queueLengths := []QueueLengths{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for qL := range queueLengthsCh {
			queueLengths = append(queueLengths, qL)
		}
	}()

	kits := []hilberthotel.WelcomeKit{}
	for busKits := range hilbertCh {
		kits = append(kits, busKits...)
	}

	wg.Wait()

	fmt.Println()
	fmt.Printf("%v guests have been given a room by Hilber at his Hotel\n", len(kits))

	return kits, queueLengths
}
