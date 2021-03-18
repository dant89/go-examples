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
	fmt.Printf("I don't know why you say, \"%s\", I say, \"%s, %s, %s\"\n", bar, foo, foo, foo)
	fmt.Println("I don't know why you say, \"" + bar + "\", I say, \"" + foo + "\"")
}
