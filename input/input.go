package main

import (
	"fmt"
	"html"
	"strconv"
)

func main()  {
	var beerCount int

	var rawBeerEmoji int = 127866
	var beerEmoji string = html.UnescapeString("&#" + strconv.Itoa(rawBeerEmoji) + ";");
	fmt.Printf("How many %s would you like?\n", beerEmoji)

	fmt.Scan(&beerCount)
	fmt.Printf("You have chosen %d %s's!\n", beerCount, beerEmoji)
}
