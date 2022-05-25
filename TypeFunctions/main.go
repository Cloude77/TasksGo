package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func add1(x int, y int) int {
	return x + y
}
func multiply(x int, y int) int {
	return x * y
}

func display(message string) {
	fmt.Println(message)
}

func main() {
	var f func(int, int) int = add
	fmt.Println(f(3, 4))

	var x = f(4, 5)
	fmt.Println(x)

	f = add1
	fmt.Println(f(3, 4))

	f = multiply
	fmt.Println(f(3, 4))

	var f1 func(string) = display
	f1("hello")

}
