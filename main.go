package main

import "fmt"

func canIDrink(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 50:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(15))
}
