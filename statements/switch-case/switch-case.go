package main

import "fmt"

func main()  {
	
	for i := 1; i <= 5; i++ {

		// Tagged switch, i = tag
		switch i {
			case 1:
				fmt.Printf("Case 1 equals: %d\n", i)
			case 2:
				fmt.Printf("Case 2 equals: %d\n", i)
			case 3:
				fmt.Printf("Case 3 equals: %d\n", i)
			default:
				fmt.Printf("Case not found, handling via default: %d\n", i)
		}
	}

	for i := 1; i <= 5; i++ {

		// Tagless switch
		// First true matching condition is executed
		switch {
			case i > 2 && i <= 3:
				fmt.Printf("Case 1 equals: %d\n", i)
			case i < 2:
				fmt.Printf("Case 2 equals: %d\n", i)
			case i > 3:
				fmt.Printf("Case 3 equals: %d\n", i)
			default:
				fmt.Printf("Case not found, handling via default: %d\n", i)
		}

	}
}
