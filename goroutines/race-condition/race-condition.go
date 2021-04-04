package main

import (
	"fmt"
	"time"
)

var i int

func main() {
	go subOne()
	go subTwo()

	/*
	When the main function ends, all goroutines end.
	
	We can do a hacky fix using sleep but this should
	be avoided for multiple reasons. One example is that
	timings may change within a goroutine and then it
	would be ended prior to completing. Another example
	is that the operating system may prioritise another
	thread or goroutine first.

	Timing is nondeterministic.

	Debugging errors that only happen randomly, say 1/10
	due to timings issue is very challenging so we want to
	avoid that at all costs.
	
	Other examples will show better soloutions to this issue.
	*/
	time.Sleep(1 * time.Second)
}

func subOne() {
	i++
	fmt.Println("This is goroutine 1, sample int value:", i)
}

func subTwo() {
	i--
	fmt.Println("This is goroutine 2, sample int value:", i)
}
