package main

import "fmt"

func main()  {
	
	// Deferred function calls, run at the end of the surrounding function
	// Arguments to these functions are evaluated immediately
	defer fmt.Println("These function examples were great, goodbye!")

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

	funcVar = incFn
	fmt.Println("Variable as a function example output:", funcVar(1))

	var sampleOutput int
	sampleOutput = applyInt(funcVar, 1)
	fmt.Println("Function as a parameter example output:", sampleOutput)

	var sampleOutputTwo int
	sampleOutputTwo = applyInt(func (x int) int {
		return x + 3
	}, 1)
	fmt.Println("Annonymous function as a parameter example output:", sampleOutputTwo)
	
	maxValue := getMax(2, 4, 99, 2, 4, 2, 423)
	fmt.Println("The max value from multiple argument input is:", maxValue)

	maxValueSlice := []int{33, 63, 1, 2, 2323, 33}
	fmt.Println("The max value from slice argument input is:", getMax(maxValueSlice...))
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

// Example of setting a variable as a function
var funcVar func(int) int

func incFn(x int) int {
	return x + 1
}

// Example function passing func as a parameter
func applyInt(afunct func (int) int, val int) int {
	val = val + 1
	return afunct(val)
}

// Variale Argument Number
func getMax(vals ...int) int {
	maxValue := -1
	for _, v := range vals {
		if v > maxValue {
			maxValue = v
		}
	}
	return  maxValue
}

