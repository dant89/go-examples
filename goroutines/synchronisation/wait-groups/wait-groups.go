package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go subOne(&wg, 1)
	go subOne(&wg, 2)
	go subOne(&wg, 3)
	wg.Wait()
	fmt.Println("The Goroutines are completed, we can continue!")
}

func subOne(wg *sync.WaitGroup, i int) {
	fmt.Printf("Goroutine %d completed\n", i)
	wg.Done()
}
