package main

import "fmt"

var x int = 5

func main() {
	scopeTestA()
	scopeTestB()
	x = 7
	scopeTestA()
}

func scopeTestA() {
	fmt.Println("Variable x equals", x)
}

func scopeTestB() {
	x = 3
	fmt.Println("Variable x equals", x)
}
