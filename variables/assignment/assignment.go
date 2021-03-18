package main

import (
	"fmt"
	"html"
	"strconv"
)

func main() {

	var x int = 3
	var y int8 = 127
	var z uint64 = 4294967295

	var dbId string = "30594050-c1bd-ec25-b474-5360e7a700dd"

	var piSmall float32 = 3.141592 // ~6 digits of precision
	var piLarge float64 = 3.141592653589793 // ~15 digits of precision

	var rawSmileEmoji int = 128522

	fmt.Printf("%T %d\n", x, x)
	fmt.Printf("%T %d\n", y, y)
	fmt.Printf("%T %d\n", z, z)

	fmt.Printf("%T %s\n", dbId, dbId)

	fmt.Printf("%T %.6f\n", piSmall, piSmall)
	fmt.Printf("%T %.15f\n", piLarge, piLarge)

	// An example of converting the smile dec emoji value into html character
	var smileEmoji string = html.UnescapeString("&#" + strconv.Itoa(rawSmileEmoji) + ";");
	fmt.Printf("%s\n", smileEmoji)
}
