package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)

	for i := 1; i <= 10; i++ {
		go generator(c, i)
	}

	for v := range c {
		go processor(v)
	}
}

func generator(c chan int, i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Generating...", i)
	c <- i
	close(c)
}

func processor(v int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Processing...", v)
}
