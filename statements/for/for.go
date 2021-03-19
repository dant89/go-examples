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
	for m > 1 {
		fmt.Printf("For m equals: %d\n", m)
		m--
	}

	var n int = 1
	for n < 10 {
		if (n > 5) {
			// if the count goes over 5, break and exit the for loop
			break;
		}
		fmt.Printf("For n equals: %d\n", n)
		n++
	}

	var o int = 1
	for o < 10 {
		o++
		if (o == 5) {
			// if the count equals 5, skip this for loop iteration
			continue
		}
		fmt.Printf("For o equals: %d\n", o)
	}	

	// This is possible but has to be used carefully and in the right scenario
	//for {
	//	fmt.Println("Infinte loop.")
	//}
}
