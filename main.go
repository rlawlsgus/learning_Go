package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("kim", "park", "li")
}