package main

import (
    "fmt"
)

func main() {
    var numberOfCustomers int
    fmt.Print("Enter the number of customers: ")
    fmt.Scanln(&numberOfCustomers)

    for i := 0; i < numberOfCustomers; i++ {
        var accountNumber, initialBalance, totalCharges, totalCredits, creditLimit int
        
        fmt.Print("Enter account number: ")
        fmt.Scanln(&accountNumber)
        fmt.Print("Enter balance at beginning of month: ")
        fmt.Scanln(&initialBalance)
        fmt.Print("Enter total amount of all items: ")
        fmt.Scanln(&totalCharges)
        fmt.Print("Enter total credits of customer: ")
        fmt.Scanln(&totalCredits)
        fmt.Print("Enter customer credit limit: ")
        fmt.Scanln(&creditLimit)
        
        newBalance := initialBalance + totalCharges - totalCredits

        fmt.Printf("Customer's new balance is: %d\n", newBalance)

        if newBalance > creditLimit {
            fmt.Println("Credit limit exceeded")
        } else {
            fmt.Println("Credit limit not exceeded")
        }
    }
}