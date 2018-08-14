package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*
Going to pick a random number between 1-100
have the user try to guess in 5 tries or less
*/

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	myNum := rand.Intn(100)

	fmt.Println("I have chosen a number between 1 and 100\n")
	fmt.Println("You have 5 guesses\n")

	for i := 1; i < 6; i++ {
		var guess int
		fmt.Println("Guess #", i)
		fmt.Scanln(&guess)
		// ToDo assert user entered a number
		if guess > myNum {
			fmt.Println("lower")
		} else if guess < myNum {
			fmt.Println("Higher")
		} else {
			fmt.Println("Congratulations, you are correct")
			os.Exit(0)
		}
	}

	fmt.Println("Sorry, you did not guess my number: ", myNum)
	os.Exit(1)
}
