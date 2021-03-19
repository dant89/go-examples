package main

import "fmt"

func main() {
	pointerExample()
	newPointerExample()
}

func pointerExample() {
	var x int = 23
	var y int

	var z *int // pointer to an int

	z = &x // z now points to x
	y = *z // z is now 1

	fmt.Println("x", x, "equals y", y)
}

func newPointerExample() {
	ptr := new(int)
	*ptr = 5

	fmt.Println("Pointer memory location equals", &ptr)
	fmt.Println("Pointer equals", *ptr)
}
