package main

import "fmt"

func main() {
	count := 1
	for count <= 3 {
		var name string
		var earnings float64
		var tax float64

		fmt.Println("Enter citizen", count, "'s name:")
		fmt.Scanln(&name)

		fmt.Println("Enter", name, "'s earnings for the year:")
		fmt.Scanln(&earnings)

		if earnings <= 30000 {
			tax = 0.15 * earnings
		} else {
			tax = (0.15 * 30000) + (0.20 * (earnings - 30000))
		}

		fmt.Println(name, "'s total tax: $", tax)
		count++
	}
}