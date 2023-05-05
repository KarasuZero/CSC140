package main

import (
	"fmt"
)

func main() {

	var guess int
	var tries int

	// main loop
	for {
		// get guess from user
		fmt.Print("Please enter your guess: ")
		fmt.Scanln(&guess)
		fmt.Println("You guessed: ", guess)
		// check if guess is a space
		if guess == 0 {
			fmt.Println("You guessed a space!")
			break
		}
		tries++
	}
}
