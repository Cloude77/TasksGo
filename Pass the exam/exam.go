package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade")
	reader := bufio.NewReader(os.Stdin)   // keyboard reading
	input, err := reader.ReadString('\n') // pressing enter
	if err != nil {
		log.Fatal(err) // abort the program if error
	}

	input = strings.TrimSpace(input)            // delete symbol
	grade, err := strconv.ParseFloat(input, 64) // string in float 64
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade > 60 {
		status = "passing"
	} else {
		status = "failing"

	}
	fmt.Println("A grade of", grade, "is", status)
}
