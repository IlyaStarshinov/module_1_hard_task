package fan_in

import (
	"log"
	"sync"
)

// MergeChannels - принимает несколько каналов на вход и объединяет их в один
// Fan-in и merge channels синонимы
func MergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int, workerID int) {
		defer wg.Done()
		for val := range ch {
			log.Printf("Gorutina №%d processed the value of %d", workerID, val)
			out <- val
		}
	}

	wg.Add(len(channels))
	for id, ch := range channels {
		go output(ch, id)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
