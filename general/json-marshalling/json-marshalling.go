package main

import (
	"encoding/json" 
	"fmt"
	"os"
)

type Car struct {
	Name string
	Colour string
	Doors int8
	TopSpeed int16
}

func main()  {
	
	car := Car{Doors: 4, TopSpeed: 140, Colour: "black", Name: "Range Rover"}
	barr, err := json.Marshal(car)
	if err != nil {
		fmt.Println("An error was encountered whilst trying to encode to JSON.")
	} else {
		os.Stdout.Write(barr)
	}

	var car2 Car
	err2 := json.Unmarshal(barr, &car2)
	if err2 != nil {
		fmt.Println("An error was encountered whilst trying to decode JSON.")
	}
	fmt.Println(car2)
}
