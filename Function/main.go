package main

import "fmt"

func main() {
	hello()
	hello()
	hello()

	add(4, 5)
	add(20, 6)

	add1(1, 2, 3.4, 5.6, 1.2)

	add2(1, 2, 3)
	add2(1, 2, 3, 4)
	add2(1, 2, 3, 4, 5, 6, 7, 8)

}

func hello() {
	fmt.Println("Hello my friend!")

}

func add(x int, y int) {
	var z = x + y
	fmt.Println("x + y = ", z)

}

func add1(x, y int, a, b, c float64) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}
func add2(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}
