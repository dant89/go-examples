package main

import (
	"fmt"
	"strings"
)

// https://golang.org/pkg/strings/
func main()  {
	var myString string = "This is some normal test in a string"
	fmt.Println(myString)

	myString = strings.ReplaceAll(myString, "normal", "amazing")
	fmt.Println(myString)

	fmt.Println(strings.Title(myString))

	fmt.Println(strings.ToUpper(myString))
}
