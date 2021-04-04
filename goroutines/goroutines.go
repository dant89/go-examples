package main

import (
	"fmt"
	"time"
)

func main() {
	go subOne()
	go subTwo()

	// When the main function ends, all goroutines end.
	//
	// We can do a hacky fix using sleep but this should
	// be avoided for multiple reasons. One example is that
	// timings may change within a goroutine and then it
	// would be ended prior to completing. Other examples
	// will show better soloutions to this issue.
	time.Sleep(1 * time.Second)
}

func subOne() {
	fmt.Println("This is goroutine 1")
}

func subTwo() {
	fmt.Println("This is goroutine 2")
}
