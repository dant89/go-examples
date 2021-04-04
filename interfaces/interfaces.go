package main

import "fmt"

type Wheel interface {
	WheelRadius() int
}

type Car struct {
	Wheelsize int
	Colour string
}

type Van struct {
	Wheelsize int
	Colour string
	StorageSpace float32
}

func (c Car) WheelRadius() int {
	return c.Wheelsize * 3
}

func (v Van) WheelRadius() int {
	return v.Wheelsize * 5
}

func WheelExploded(w Wheel) bool {
	return true
}

func main() {
	myCar := Car{4, "red"}
	fmt.Printf("My cars wheel radius is: %d\n", myCar.WheelRadius())

	myVan := Van{4, "white", 5.66}
	fmt.Printf("My vans wheel radius is: %d\n", myVan.WheelRadius())

	var myWheel Wheel
	myWheel = myCar
	
	if WheelExploded(myWheel) {
		fmt.Println("My car wheel has exploded!!")
	}
}
