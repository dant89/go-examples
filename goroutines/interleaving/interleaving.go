package main

import (
	"fmt"
	"sync"
)

var i int

func main() {
	
	var wg sync.WaitGroup

	for l := 0; l < 100000; l++ {
		i = 0

		wg.Add(2)
		go subOne(&wg)
		go subOne(&wg)
		wg.Wait()
		
		if (i == 1) {
			fmt.Printf("Here is an example where i was read before being updated in both instances, attempt number %d\n", l)
		}
	}
}

func subOne(wg *sync.WaitGroup) {
	i++
	wg.Done()
}
