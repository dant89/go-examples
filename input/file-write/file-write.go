package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data := []byte("Hello, world!")

	err := ioutil.WriteFile("test-output.txt", data, 777)
	if err != nil {
		fmt.Printf("A write error occurred!")
	}
}
