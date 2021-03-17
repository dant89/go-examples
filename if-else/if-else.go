package main

import "fmt"

func main() {

	if 23%2 == 0 {
		fmt.Println("23 is even")
	} else {
		fmt.Println("23 is odd")
	}

	num := 123
	numMax := 124
	if num < numMax {
		fmt.Println(num, "is smaller than", numMax)
	}
}
