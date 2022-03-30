package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(reflect.TypeOf(7.5))
	// fmt.Println(reflect.TypeOf("string"))

	// 	quantity := 4
	// 	lenght, width := 1.2, 2.4
	// 	customerName := "Damon Cole"

	// 	fmt.Println(customerName)
	// 	fmt.Println("has ordered", quantity, "sheets")
	// 	fmt.Println("each with an area of")
	// 	fmt.Println(lenght*width, "square meters")
	// var myInt int = 2
	// fmt.Println(reflect.TypeOf(myInt))
	// fmt.Println(reflect.TypeOf(float64(myInt)))

	var len float64 = 1.2
	var width int = 2
	fmt.Println("Area is", len*float64(width))
	fmt.Println("lenght > width", len > float64(width))

	var now time.Time = time.Now()
	var year int = now.Year() // Year - method
	fmt.Println(year)

	for x := 1; x >= 3; x++ {
		fmt.Println(x)
	}

}
