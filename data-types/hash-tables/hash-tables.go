package main

import "fmt"

func main()  {
	
	m := make(map[string]string)

	m["bob"] = "1 street"
	m["brenda"] = "2 street"
	m["geff"] = "3 street"

	fmt.Println("The address book has", len(m), "people in!")

	for key, value := range m {
		fmt.Println(key, "lives at", value)
	}
}
