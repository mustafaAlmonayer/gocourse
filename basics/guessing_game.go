package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingGame() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Try to guess a number between 1 and 100.")

	for {
		var guess int
		fmt.Println("Guess the number: ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You've guessed the number!")
			break
		}
	}

}
