package main

import "fmt"

func add(x, y int) int      { return x + y }
func subtract(x, y int) int { return x - y }
func multiply(x, y int) int { return x * y }

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return add
	} else if n == 2 {
		return subtract
	} else {
		return multiply
	}
}

func main() {

	f := selectFn(1)
	fmt.Println(f(3, 4))

	f = selectFn(3)
	fmt.Println(f(3, 4))

}
