package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	target := generateNumber()
	fmt.Println("Your Number is: ", target)

	var guess int
	var tries int

	// main loop
	for {
		// get guess from user
		fmt.Println("Please enter your guess: ")
		scan, err := fmt.Scan(&guess)
		if err != nil {
			fmt.Print(scan)
			return
		}

		// check if guess are not empty
		if guess == 0 {
			fmt.Println("Invalid input! Please enter an integer.")
			continue
		}

		// check if guess are int
		// I grabbed this int validation from w3 schools, I'm not sure how to check if a obj is an int in go
		if _, err := fmt.Scan(&guess); err != nil {
			fmt.Println("Invalid input! Please enter an integer.")
			continue
		}

		// check if guess are between 1 and 100
		if !validateGuess(guess) {
			fmt.Println("Invalid guess! Please enter an integer between 1 and 100.")
			continue
		}

		// check guesses
		tries++

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You guessed it!")
			fmt.Println("It took you ", tries, " tries.")
			break
		}
	}
}

func generateNumber() int {
	// generate a random number between 1-100 using current time as seed
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	return target
}

func validateGuess(guess int) bool {
	if guess < 1 || guess > 100 {
		fmt.Println("Your guess must be between 1 and 100.")
		return false
	}
	return true
}

func gameLoop() {

}
