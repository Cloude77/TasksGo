package main

import "fmt"

func main() {

	//	numbers := [5]int{1, 2, 3, 4, 5}

	var numbers1 [5]int = [5]int{1, 2}
	fmt.Println(numbers1)

	var numbers = [...]int{1, 2, 3, 4, 5} // len 5
	numbers2 := [...]int{1, 2, 3}         // len 3
	fmt.Println(numbers)
	fmt.Println(numbers2)

	//	var values [5]int = [5]int{1, 2, 3, 4, 5}
	values := [5]int{1, 2, 3, 4, 5}

	fmt.Println(values[2])
	fmt.Println(values[4])
	values[0] = 90
	fmt.Println(values[0])

	// index

	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2])

}
