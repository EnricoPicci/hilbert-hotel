package hilberthotelconcurrentrecursive

import (
	"sort"
	"testing"
	"time"

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
	delay := 10 * time.Microsecond

	keysCh := make(chan int)
	go RoomKeysClerk(11, keysCh)

	hilbertCh := make(chan []hilberthotel.WelcomeKit)
	go BusClerk(1, keysCh, hilbertCh, buffer, delay)

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

func TestHilbertMassive(t *testing.T) {
	buffer := 100
	delay := 10 * time.Microsecond
	numOfPassengers := 100000

	kits := Hilbert(numOfPassengers, buffer, delay, true)

	if len(kits) != numOfPassengers {
		t.Errorf("Created %v kits ==> expected %v", len(kits), numOfPassengers)
	}
}
