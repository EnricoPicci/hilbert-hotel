package hilberthotelnonrecursive

import (
	"sort"
	"testing"
	"time"

	"github.com/EnricoPicci/hilberthotel"
)

func TestWelcomeKitsForBus(t *testing.T) {
	delay := 10 * time.Microsecond

	Kits := WelcomeKits(11, delay)

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
	gotNumOfKits := len(Kits)
	if gotNumOfKits != expectedNumOfKits {
		t.Errorf("expected %v, got %v", expectedNumOfKits, gotNumOfKits)
	}

	sort.Slice(Kits, func(i, j int) bool {
		if Kits[i].BusNumber == Kits[j].BusNumber {
			return Kits[i].PassengerNumber < Kits[j].PassengerNumber
		}
		return Kits[i].BusNumber < Kits[j].BusNumber
	})

	for i, gotEnvelop := range Kits {
		expectedEnvelop := expectedKits[i]
		if gotEnvelop.RoomNumber != expectedEnvelop.RoomNumber {
			t.Errorf("Room number in Kit %v ==> expected %v - got %v", i, expectedEnvelop, gotEnvelop)
		}
	}
}

func TestHilbertMassive(t *testing.T) {
	delay := 10 * time.Microsecond
	numOfPassengers := 10000

	kits := Hilbert(numOfPassengers, delay, true)

	if len(kits) != numOfPassengers {
		t.Errorf("Created %v kits ==> expected %v", len(kits), numOfPassengers)
	}
}
