package hilberthotelnonrecursive

import (
	"fmt"
	"time"

	"github.com/EnricoPicci/hilberthotel"
)

func NewWelcomeKit(busNumber int, passengerNumber int, row []int, delay time.Duration) hilberthotel.WelcomeKit {
	return hilberthotel.NewWelcomeKit(busNumber, passengerNumber, row[len(row)-busNumber], delay)
}

func WelcomeKits(upTo int, delay time.Duration) []hilberthotel.WelcomeKit {
	rows := [][]int{}
	counter := 0
	i := 0
	for {
		counter++
		var row []int
		for j := 0; j < counter; j++ {
			if i+j == upTo {
				if row != nil {
					rows = append(rows, row)
				}
				goto rowsReady
			}
			if row == nil {
				row = make([]int, counter)
			}
			row[j] = i + j + 1
		}
		rows = append(rows, row)
		i = i + counter
	}
rowsReady:

	welcomeKits := []hilberthotel.WelcomeKit{}
	rowNumber := 1
	var passengerNumbersForBus []int

	for _, row := range rows {
		passengerNumbersForBus = append(passengerNumbersForBus, 1)
		for busNumber := 0; busNumber < rowNumber; busNumber++ {
			passengerNumber := passengerNumbersForBus[busNumber]
			welcomeKit := NewWelcomeKit(busNumber+1, passengerNumber, row, delay)
			if welcomeKit.RoomNumber > 0 {
				welcomeKits = append(welcomeKits, welcomeKit)
			}
			passengerNumbersForBus[busNumber]++
		}
		rowNumber++
	}

	return welcomeKits
}

func Hilbert(upTo int, delay time.Duration, verbose bool) []hilberthotel.WelcomeKit {
	kits := WelcomeKits(upTo, delay)

	if verbose {
		fmt.Println()
		fmt.Printf("%v guests have been given a room by Hilber at his Hotel", len(kits))
	}

	return kits
}
