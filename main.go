package main

import "fmt"

func main() {
	kim := map[string]string{"name": "kim", "age": "12"}
	for _, value := range kim {
		fmt.Println(value)
	}
}