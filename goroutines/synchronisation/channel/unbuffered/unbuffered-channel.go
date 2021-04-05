package main

import "fmt"

func main() {
	// A channel is unbuffered by default which means it cannot hold data in transit
	c := make(chan int)

	go multiply(2, 4, c)
	go multiply(4, 6, c)

	a := <- c
	b := <- c

	fmt.Println("The summed values total: ", a*b)
}

func multiply(int1 int, int2 int, c chan int) {
	c <- int1 * int2
}
