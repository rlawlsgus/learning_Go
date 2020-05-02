package main

import "fmt"

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"gukbap", "dupbap"}
	kim := person{name: "kim", age: 17, favFood: favFood}
	fmt.Println(kim)
}