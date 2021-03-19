package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	
	var x [5]int
	x[2] = 2

	for _, v := range x {
		fmt.Println(v)
	}

	// Array literals
	var daysOfWeek [7]string = [7]string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}
	for i, v := range daysOfWeek {
		if (v == "Friday") {
			fmt.Printf("Day %d is %s, wooo!\n", i, v)
		}
	}

	// Slice
	var midweekDays string
	midweekDaysSlice := daysOfWeek[0:5]
	for _, v := range midweekDaysSlice {
		if (v == "Friday") {
			midweekDays = midweekDays + " and " + v
		} else if midweekDays != "" {
			midweekDays = midweekDays + ", " + v
		} else {
			midweekDays = v
		}
	}

	fmt.Println(midweekDays, "are the days in the middle of the week")
	fmt.Printf("There are %d days in the middle of the week\n", len(midweekDaysSlice))

	// Slice dice example
	diceSlice := make([]int, 6)
	diceSlice[0] = 1
	diceSlice[1] = 2
	diceSlice[2] = 3
	diceSlice[3] = 4
	diceSlice[4] = 5
	diceSlice[5] = 6

	rand.Seed(time.Now().UnixNano())
	var randomDiceInt int = rand.Intn(6)
	fmt.Printf("You rolled a %d on the dice\n", diceSlice[randomDiceInt])
}
