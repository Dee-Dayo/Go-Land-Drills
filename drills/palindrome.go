package main

import (
	"fmt"
)

func main() {
	var userInput int
	fmt.Print("Enter a five digits number: ")
	fmt.Scanln(&userInput)

	for userInput != -1 {
		if userInput < 10000 || userInput > 99999 {
			fmt.Print("Enter a valid number: ")
			fmt.Println()
		} else {
			userInput1 := userInput / 10000
			userInput2 := (userInput / 1000) % 10
			userInput3 := (userInput / 100) % 10
			userInput4 := (userInput / 10) % 10
			userInput5 := userInput % 10

			userInput6 := (userInput5 * 10000) + (userInput4 * 1000) + (userInput3 * 100) + (userInput2 * 10) + userInput1

			if userInput == userInput6 {
				fmt.Printf("%d is a palindrome.\n", userInput)
				fmt.Println()
			} else {
				fmt.Printf("%d is not a palindrome number.\n", userInput)
				fmt.Println()
			}
		}
		fmt.Print("Enter a five digits number: ")
		fmt.Scanln(&userInput)
	}

	if userInput == -1 {
		fmt.Println("Bye")
	}
}