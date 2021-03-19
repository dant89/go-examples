package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// content equals []byte
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("A read error occurred!")
	}

	fmt.Printf("File contents: %s\n", content)

	f, error := os.Open("test.txt")
	if error != nil {
		fmt.Println("Could not open file")
	}
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	if error != nil {
		fmt.Println("Could not open file")
	}
	fmt.Printf("Read %d bytes\n", nb)
	f.Close()
}
