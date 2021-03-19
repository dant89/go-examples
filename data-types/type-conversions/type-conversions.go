package main

import "fmt"

func main() {
	var x int32 = 6
	var y int64 = 12

	x = int32(y)
	fmt.Printf("32bit x equals conversion of 64bit y: %d\n", x)

	var m int8 = 4
	var n int32 = 2147483647

	fmt.Printf("8bit m equals 8bit: %d\n", m)
	m = int8(n)
	// this returns -1 because n value does not fit 8bit
	fmt.Printf("8bit m equals conversion of 32bit n: %d\n", m)
}
