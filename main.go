package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 20
	// why not b := *&a ????????
	fmt.Println(a)
}