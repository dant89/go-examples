package main

import "fmt"

func main()  {
	
	var x int = 1
	var y int = 1

	getOutput(x, "x")
	getOutput(y, "y")

	x = getExampleOne(x)
	getOutput(x, "x")

	x, y = getExampleTwo(x, y)
	getOutput(x, "x")
	getOutput(y, "y")

	callByReferenceExample(&x)
	getOutput(x, "x")
}

func getExampleOne(x int) int {
	return x * 3
}

func getExampleTwo(x , y int) (int, int) {
	return x, y + 1
}

func getOutput(x int, name string) {
	fmt.Printf("%s equals %d\n", name, x)
}

func callByReferenceExample(y *int) {
	*y = *y + 2
}
