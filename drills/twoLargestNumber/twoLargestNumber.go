package main

import "fmt"

func main() {
	counter := 1
	largest := 0
	secondLargest := 0

	for counter <= 10 {
		fmt.Printf("Enter number %d: ", counter)
		var num int
		fmt.Scanln(&num)
		
		if num > largest {
			secondLargest = largest
			largest = num
		} else if num > secondLargest {
			secondLargest = num
		}
		counter++
	}

	fmt.Println("The largest number is", largest)
	fmt.Println("The second largest number is", secondLargest)
}