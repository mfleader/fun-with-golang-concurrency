package main

import (
	"fmt"
	"sync"
)

func processing(item string, sem chan struct{}, wg *sync.WaitGroup) {
	sem <- struct{}{}
	defer func() {
		// read an item, then drop it
		<-sem
	}()
	// when the processing function finishes,
	// tell your waitgroup that you're all done
	defer wg.Done()
	fmt.Printf("Processing item %s...\n", item)
}

func main() {
	// fixed-capacity parallel processor

	// items := []string{"0", "1", "2", "3"}
	items := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40"}
	// make a channel with capacity 10
	sem := make(chan struct{}, 10)
	//sem := make(chan bool, 10)
	wg := &sync.WaitGroup{}

	// increment waitgroup counter by length of items
	wg.Add(len(items))

	for _, item := range items {
		// give item a local scope, so that each goroutine
		// uses a different item
		item := item

		// processing function
		go func() {
			// to start a goroutine
			// send a value of the channel's type into the channels
			// this increments a counter in the channel
			//sem <- struct{}{}
			//sem <- true

			defer func() {
				// read/take/consume an item from the channel
				// block until we receive a notification from
				// the worker on the channel
				<-sem
			}()

			// when the processing function finishes,
			// tell your waitgroup that you're all done
			// i.e. decrement waitgroup counter by 1
			defer wg.Done()

			// do the thing
			fmt.Printf("Processing item %s...\n", item)

		}()

		// go processing(item, sem, wg)

	}

	// wait for all processing to complete
	// i.e. block until the waitgroup counter is 0
	wg.Wait()

}
