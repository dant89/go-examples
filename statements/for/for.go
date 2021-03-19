package main

import "fmt"

func main()  {
	
	var count int = 10
	for i := 1; i < count; i++ {
		fmt.Printf("For i equals: %d\n", i)
	}

	for l := 20; l >= count; l-- {
		fmt.Printf("For l equals: %d\n", l)
	}

	var m int = 5
	for m < 10 {
		fmt.Printf("For m equals: %d\n", m)
		m--
	}	

	// This is possible but has to be used carefully and in the right scenario
	//for {
	//	fmt.Println("Infinte loop.")
	//}
}
