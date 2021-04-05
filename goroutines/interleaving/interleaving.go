package main

import (
	"fmt"
	"sync"
)

var i int
var mut sync.Mutex

func main() {
	interleavingExample()
	noneInterleavingExample()
}

func interleavingExample() {
	var wg sync.WaitGroup

	for l := 0; l < 100000; l++ {
		i = 0

		wg.Add(2)
		go subInterleaving(&wg)
		go subInterleaving(&wg)
		wg.Wait()
		
		if (i == 1) {
			fmt.Printf("Interleaving occured during test1 attempt %d\n", l)
		}
	}
}

func noneInterleavingExample() {
	var wg sync.WaitGroup
	var interleaved bool = false

	for l := 0; l < 100000; l++ {
		i = 0

		wg.Add(4)
		go subNoneInterleavinig(&wg)
		go subNoneInterleavinig(&wg)
		go subNoneInterleavinig(&wg)
		go subNoneInterleavinig(&wg)
		wg.Wait()
		
		if (i == 1) {
			interleaved = true
		}
	}

	if interleaved {
		fmt.Println("Interleaving occured during test2")
	} else {
		fmt.Println("Interleaving did not occur during test2")
	}
}

func subInterleaving(wg *sync.WaitGroup) {
	i++
	wg.Done()
}

func subNoneInterleavinig(wg *sync.WaitGroup) {
	mut.Lock()
	i++
	mut.Unlock()
	wg.Done()
}
