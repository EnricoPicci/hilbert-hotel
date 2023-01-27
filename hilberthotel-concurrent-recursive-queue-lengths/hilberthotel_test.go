package main

import (
	"fmt"
	"sort"
	"testing"

	"github.com/EnricoPicci/hilberthotel"
)

func TestRoomKeysClerk_10(t *testing.T) {
	keysCh := make(chan int)
	go RoomKeysClerk(10, keysCh)

	roomNumbers := []int{}

	for n := range keysCh {
		roomNumbers = append(roomNumbers, n)
	}

	var expectedNums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, r := range roomNumbers {
		if r != expectedNums[i] {
			t.Errorf("Room numbers ==> expcted %v - got %v", expectedNums[i], r)
		}
	}

}

func TestHilbertHospitality(t *testing.T) {
	buffer := 10

	keysCh := make(chan int)
	go RoomKeysClerk(11, keysCh)

	hilbertCh := make(chan []hilberthotel.WelcomeKit)
	queueLengthsCh := make(chan QueueLengths, buffer)
	go BusClerk(1, keysCh, hilbertCh, queueLengthsCh, buffer)

	kits := []hilberthotel.WelcomeKit{}
	for busKits := range hilbertCh {
		kits = append(kits, busKits...)
	}

	expectedKits := []hilberthotel.WelcomeKit{
		{BusNumber: 1, PassengerNumber: 1, RoomNumber: 1},
		{BusNumber: 1, PassengerNumber: 2, RoomNumber: 3},
		{BusNumber: 1, PassengerNumber: 3, RoomNumber: 6},
		{BusNumber: 1, PassengerNumber: 4, RoomNumber: 10},
		{BusNumber: 2, PassengerNumber: 1, RoomNumber: 2},
		{BusNumber: 2, PassengerNumber: 2, RoomNumber: 5},
		{BusNumber: 2, PassengerNumber: 3, RoomNumber: 9},
		{BusNumber: 3, PassengerNumber: 1, RoomNumber: 4},
		{BusNumber: 3, PassengerNumber: 2, RoomNumber: 8},
		{BusNumber: 4, PassengerNumber: 1, RoomNumber: 7},
		{BusNumber: 5, PassengerNumber: 1, RoomNumber: 11},
	}

	expectedNumOfKits := len(expectedKits)
	gotNumOfKits := len(kits)
	if gotNumOfKits != expectedNumOfKits {
		t.Errorf("expected %v, got %v", expectedNumOfKits, gotNumOfKits)
	}

	sort.Slice(kits, func(i, j int) bool {
		if kits[i].BusNumber == kits[j].BusNumber {
			return kits[i].PassengerNumber < kits[j].PassengerNumber
		}
		return kits[i].BusNumber < kits[j].BusNumber
	})

	for i, gotEnvelop := range kits {
		expectedEnvelop := expectedKits[i]
		if gotEnvelop.RoomNumber != expectedEnvelop.RoomNumber {
			t.Errorf("Room number in envelope %v ==> expected %v - got %v", i, expectedEnvelop, gotEnvelop)
		}
	}
}

func TestQueueLengths(t *testing.T) {
	buffer := 1
	numOfPassengers := 11

	kits, queueLengths := Hilbert(numOfPassengers, buffer)

	if len(kits) != numOfPassengers {
		t.Errorf("Created %v kits ==> expected %v", len(kits), numOfPassengers)
	}
	expectedNumOfBusses := 5
	gotNumOfBusses := len(queueLengths)
	if gotNumOfBusses != expectedNumOfBusses {
		t.Errorf("The number of queue lengths created should be %v - instead is %v", expectedNumOfBusses, gotNumOfBusses)
	}
}

func TestHilbertMassive(t *testing.T) {
	buffer := 1000
	numOfPassengers := 100001

	kits, queueLengths := Hilbert(numOfPassengers, buffer)

	if len(kits) != numOfPassengers {
		t.Errorf("Created %v kits ==> expected %v", len(kits), numOfPassengers)
	}
	if len(queueLengths) == 0 {
		t.Error("No queue lengths created")
	}

	for _, ql := range queueLengths {
		fmt.Printf("Bus %v \t - \t avgLen: %.2f \t - \t stdDevLen: %.2f \t - \t numOfKeysReceived: %v\n", ql.busNumber, ql.avgLen, ql.stdDevLen, len(ql.lengths))
	}
}

func TestHilbertWithBuffer(t *testing.T) {
	buffer := 10
	numOfPassengers := 1000

	kits, queueLengths := Hilbert(numOfPassengers, buffer)

	if len(kits) != numOfPassengers {
		t.Errorf("Created %v kits ==> expected %v", len(kits), numOfPassengers)
	}
	if len(queueLengths) == 0 {
		t.Error("No queue lengths created")
	}
}
