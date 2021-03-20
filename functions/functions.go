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

	z := []int{1, 2, 3}
	for i, l := range z {
		a := "z-1-" + fmt.Sprint(i)
		getOutput(l, a)	
	}
	callByReferenceSliceExample(z)

	for i, l := range z {
		a := "z-2-" + fmt.Sprint(i)
		getOutput(l, a)	
	}
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

// In Go it's recommended to pass by slice as the pointser are included
// if you pass an array it's a lot more tricky to update with pointers
func callByReferenceSliceExample(sli []int) {
	sli[0] = sli[0] + 1
}
