package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i * i)
	// }

	// var i = 1
	// for i < 10 {
	// 	fmt.Println(i * i)
	// 	i++
	// }

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	// range

	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}

	var users1 = [3]string{"Tom", "Alice", "Kate"}
	for i := 0; i < len(users1); i++ {
		fmt.Println(users[i])
	}

	// break, continue
	// sum of positive numbers

	var numbers2 = [10]int{1, -2, 3, -4, 5, -6, -7, 8, -9, 10}
	var sum = 0
	for _, value := range numbers2 {
		if value < 0 {
			continue // next iteration
		}
		sum += value
	}
	fmt.Println("Sum:", sum)

	// if number > 4

	var numbers3 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sum1 = 0
	for _, value := range numbers3 {
		if value > 4 {
			break
		}
		sum1 += value
	}
	fmt.Println("Sum:", sum1)
}
