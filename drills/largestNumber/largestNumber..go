package main

import "fmt"

func main() {
	counter := 1
	largest := 0

	for counter <= 10 {
		fmt.Printf("Enter number %d: ", counter)
		var number int
		fmt.Scanln(&number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Println("The largest number is:", largest)
}