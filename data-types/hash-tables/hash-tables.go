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

	id, p := m["terry"]
	if p {
		fmt.Printf("We have Terry in the address book, he lives at %s!\n", id)
	} else {
		fmt.Println("We do not have Terry in the address book :(")
	}
}
