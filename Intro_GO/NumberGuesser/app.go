package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	for {
		fmt.Print("Guess the number between 1 and 100: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You guessed it!")
			break
		}
	}
}
