package main

import (
	"fmt"
	"html"
	"strconv"
)

func main() {

	// int examples
	var x int = 3
	var y int8 = 127
	var z uint64 = 4294967295

	fmt.Printf("%T %d\n", x, x)
	fmt.Printf("%T %d\n", y, y)
	fmt.Printf("%T %d\n", z, z)

	// string examples
	var dbId string = "30594050-c1bd-ec25-b474-5360e7a700dd"
	fmt.Printf("%T %s\n", dbId, dbId)

	// float type examples
	var piSmall float32 = 3.141592 // ~6 digits of precision
	var piLarge float64 = 3.141592653589793 // ~15 digits of precision

	fmt.Printf("%T %.6f\n", piSmall, piSmall)
	fmt.Printf("%T %.15f\n", piLarge, piLarge)

	// An example of converting the smile dec emoji value into html character
	var rawSmileEmoji int = 128522
	var smileEmoji string = html.UnescapeString("&#" + strconv.Itoa(rawSmileEmoji) + ";");
	fmt.Printf("%s\n", smileEmoji)

	// constant examples
	const constOne = 30
	const  (
		constTwo = 40
		constThree = "50"
	)
	fmt.Printf("%T %d\n", constOne, constOne)
	fmt.Printf("%T %d\n", constTwo, constTwo)
	fmt.Printf("%T %s\n", constThree, constThree)

	// iota
	type MilkyWayPlanets int
	const (
		Mercury MilkyWayPlanets = iota
		Venus 	MilkyWayPlanets = iota
		Earth 	MilkyWayPlanets = iota
		Mars 	MilkyWayPlanets = iota
		Jupiter MilkyWayPlanets = iota
		Saturn 	MilkyWayPlanets = iota
		Uranus	MilkyWayPlanets = iota
		Neptune MilkyWayPlanets = iota
	)
}
