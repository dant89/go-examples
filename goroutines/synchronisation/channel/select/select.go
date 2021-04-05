package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 5)
	abort := make(chan int)

	go generator(c, 1)
	go generator(c, 2)
	go generator(c, 3)
	go generator(c, 4)
	go generator(c, 5)
	go generator(abort, 1)

	for {
		select {
			case a := <- c:
				fmt.Println("Data recieved", a)
			case <- abort:
				fmt.Println("Application stopped")
				return
		}
	}
}

func generator(c chan int, i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Generating...", i)
	c <- i
}
