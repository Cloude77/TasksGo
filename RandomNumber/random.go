// guess - игра , в которой игрок должен угадать случайное число.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // current date in integer format
	rand.Seed(seconds)           // random number generator
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target) - prompt

	reader := bufio.NewReader(os.Stdin) // reading keyboard input

	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Println("Make a guess:")
		input, err := reader.ReadString('\n') // read the data entered by the user before pressing Enter
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)  // remove newline character
		guess, err := strconv.Atoi(input) // string is converted to integer
		if err != nil {                   // if an error the program terminates
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops, Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops, Your guess was HIGH.")

		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break

		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was: ", target)
	}

}