package main

import (
	"fmt"
	"tryme/sub"
)

func main()  {
	result:=add(1,2)
	fmt.Println("the result of the sum is:",result)
	subResult:=sub.Sub(1,2)
	fmt.Println("the sub is :",subResult)
	fmt.Println("demo")
}


func testing() bool {
	return true
}

