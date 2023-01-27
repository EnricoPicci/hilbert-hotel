package hilberthotelconcurrentrecursive

func RoomKeysClerk(upTo int, keysCh chan<- int) {
	for i := 0; i < upTo; i++ {
		keysCh <- i + 1
	}
	close(keysCh)
}
