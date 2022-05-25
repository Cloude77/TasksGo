package main

import "fmt"

func main() {
	a := 8
	b := 8
	if a < b {
		fmt.Println("a less b")
	} else if a > b {
		fmt.Println("a more b")
	} else {
		fmt.Println(" a equals b")
	}

	fmt.Println("-----------------------------")

	//switch

	num := 12
	switch num {
	case 9:
		fmt.Println("a = 10")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	default:
		fmt.Println("the value of the variable a is not defined")
	}
}
