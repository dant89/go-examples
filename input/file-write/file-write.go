package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello, world!")

	err := ioutil.WriteFile("test-output.txt", data, 777)
	if err != nil {
		fmt.Printf("A write file error occurred!")
	}

	f, err := os.Create("test-output2.txt")
	if err != nil {
		fmt.Printf("A create file error occurred!")
	}
	_, error := f.WriteString("Hello world write!")
	if error != nil {
		fmt.Printf("A write file error occurred!")
	}
}
