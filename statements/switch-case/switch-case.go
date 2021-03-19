package main

import "fmt"

func main()  {
	
	for i := 1; i <= 5; i++ {

		switch i {
			case 1:
				fmt.Printf("Case one equals: %d\n", i)
			case 2:
				fmt.Printf("Case two equals: %d\n", i)
			case 3:
				fmt.Printf("Case three equals: %d\n", i)
			default:
				fmt.Println("Case not found, handling via default")
		}

	}
}
