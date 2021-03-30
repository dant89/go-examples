package main

import "fmt"

type MyString string

func (ms MyString) Prepend (inputString string) string {
	value := inputString + string(ms)
	return value
}

func main()  {
	v := MyString("bar")
	fmt.Println(v.Prepend("foo"))
}
