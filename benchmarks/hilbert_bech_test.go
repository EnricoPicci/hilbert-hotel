package benchmarks

import (
	"flag"
	"testing"
	"time"

	"github.com/EnricoPicci/hilberthotel"
	closurerecursive "github.com/EnricoPicci/hilberthotel/hilberthotel-closure-recursive"
	concurrencyrecursive "github.com/EnricoPicci/hilberthotel/hilberthotel-concurrent-recursive"
	nonrecursive "github.com/EnricoPicci/hilberthotel/hilberthotel-nonrecursive"
)

var numOfGuests int
var delayMicrosec int

var kits []hilberthotel.WelcomeKit

func init() {
	flag.IntVar(&numOfGuests, "numGuests", 1000, "Number of guests that want to stay at Hilbert's Hotel")
	flag.IntVar(&delayMicrosec, "delayMicrosec", 10, "Delay in microsecs to make a welcome kit (simulates work to be done for each guest)")
}

func delay() time.Duration {
	return time.Duration(delayMicrosec) * time.Microsecond
}

func BenchmarkNonRecursive(b *testing.B) {
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = nonrecursive.Hilbert(numOfGuests, delay(), false)
	}
	kits = _kits
}

func BenchmarkClosureRecursive(b *testing.B) {
	var _kits []hilberthotel.WelcomeKit
	for i := 0; i < b.N; i++ {
		_kits = closurerecursive.Hilbert(numOfGuests, delay(), false)
	}
	kits = _kits
}

func BenchmarkConcurrencyRecursive(b *testing.B) {
	benchmarks := []struct {
		name   string
		buffer int
	}{
		{"buffer_0", 0},
		{"buffer_10", 10},
		{"buffer_100", 100},
		{"buffer_1000", 1000},
		{"buffer_10000", 10000},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			var _kits []hilberthotel.WelcomeKit
			for i := 0; i < b.N; i++ {
				_kits = concurrencyrecursive.Hilbert(numOfGuests, benchmark.buffer, delay(), false)
			}
			kits = _kits
		})
	}
}
