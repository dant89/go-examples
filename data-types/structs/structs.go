package main

import "fmt"

type Car struct {
	name string
	colour string
	doors int8
	topSpeed int16
}

func main()  {

	var carOne Car
	carOne.name = "Ferrari"
	carOne.colour = "Red"
	carOne.doors = 2
	carOne.topSpeed = 180

	carTwo := Car{"Lambourghini", "Orange", 2, 170}

	carThree := Car{doors: 4, topSpeed: 140, colour: "black", name: "Range Rover"}

	fmt.Printf("The %s is %s, has %d doors and can reach %dmph\n", carOne.name, carOne.colour, carOne.doors, carOne.topSpeed)
	fmt.Printf("The %s is %s, has %d doors and can reach %dmph\n", carTwo.name, carTwo.colour, carTwo.doors, carTwo.topSpeed)
	fmt.Printf("The %s is %s, has %d doors and can reach %dmph\n", carThree.name, carThree.colour, carThree.doors, carThree.topSpeed)
}
