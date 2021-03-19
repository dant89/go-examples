package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// content equals []byte
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("A read error occurred!")
	}

	fmt.Printf("File contents: %s", content)
}
