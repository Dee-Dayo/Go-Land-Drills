package main

import (
	"fmt"
)

func main() {
	tripCount := 0
	totalMiles := 0
	totalGallons := 0

	fmt.Println("Enter miles driven and gallons used for each trip (enter -1 to stop):")

	for {
		var miles, gallons int
		fmt.Print("How many Miles driven: ")
		fmt.Scanln(&miles)
		if miles == -1 {
			break
		}

		fmt.Print("How many Gallons used: ")
		fmt.Scanln(&gallons)

		tripCount++
		totalMiles += miles
		totalGallons += gallons

		milesPerGallon := float64(miles) / float64(gallons)
		fmt.Printf("Miles per gallon for trip %d: %.2f\n", tripCount, milesPerGallon)

		totalMilesPerGallon := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined miles per gallon up to trip %d: %.2f\n", tripCount, totalMilesPerGallon)
	}
}
