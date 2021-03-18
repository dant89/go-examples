package main

import (
	"fmt"
	"strings"
)

func main() {

	var foo string = "Hello"
	var bar string = "Goodbye"

	var fooLower = strings.ToLower(foo)

	fmt.Printf("You say, \"" + bar + "\", and I say, \"" + foo + ", " + fooLower + ", " + fooLower + "\"\n")

	fmt.Println("I don't know why you say, \"" + bar + "\", I say, \"" + foo + ", " + fooLower + ", " + fooLower + "\"")
	fmt.Println("I don't know why you say, \"" + bar + "\", I say, \"" + foo + "\"")

}
