package main

import "fmt"

func main() {

	const itemOne = 239.99
	const itemTwo = 129.75
	const itemThree = 99.95
	const itemFour = 350.89
	var totalValue float64

	for {
		var ItemNumber int
		fmt.Println("Enter item number or -1 to end: ")
		fmt.Scanln(&ItemNumber)
		if ItemNumber == -1 {
			break
		} else if ItemNumber < 1 || ItemNumber > 4 {
			fmt.Println("Invalid number")
			continue
		} else {
			if ItemNumber == 1 {
				totalValue += itemOne
			} else if ItemNumber == 2 {
				totalValue += itemTwo
			} else if ItemNumber == 3 {
				totalValue += itemThree
			} else if ItemNumber == 4 {
				totalValue += itemFour
			}

		}

		fmt.Println("Total Value of the Items Selected is :", totalValue)

		var salesPersonEarning = 200 + (0.09 * totalValue)
		fmt.Println("Total sales earning is :", salesPersonEarning)

	}
}