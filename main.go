package main

import "fmt"

func main() {
	names := []string{"kim", "nico", "lynn"}
	names = append(names, "flynn")
	fmt.Println(names)
}