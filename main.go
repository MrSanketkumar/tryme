package main

import (
	"fmt"
	"tryme/sub"
)

func main() {
	result := add(10, 0)
	fmt.Println("the result of the sum is:", result)
	subResult := sub.Sub(1, 2)
	fmt.Println("the sub is :", subResult)
	fmt.Println("demo")
}

func add(a, b int) int {
	if validation(a, b) {
		return a + b
	}
	return 0
}

func testing() bool {
	return true
}

func validation(a, b int) bool {
	if a > 0 && b > 0 {
		return true
	}
	return false
}
