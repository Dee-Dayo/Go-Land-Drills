package main

import (
	"fmt"
	"strings"
	"time"
)

var(
	customerName string
	cashierName string
	itemsBought []string
	piecesBought []int64
	pricePerUnit []float64
	totalPerUnit []float64
	discount float64
	currentDate = time.Now()
	subTotal float64
	discountPrice float64
	vat float64
	billTotal float64
)

func main(){
	fmt.Println("What is the customer's Name?")
	fmt.Scanln(&customerName)

	for{
		fmt.Println("What did the user buy?")
		var item string
		fmt.Scanln(&item)
		itemsBought = append(itemsBought, item)

		fmt.Println("How many pieces?")
		var piece int64
		fmt.Scanln(&piece)
		piecesBought = append(piecesBought, piece)

		fmt.Println("How much per unit?")
		var price float64
		fmt.Scanln(&price)
		pricePerUnit = append(pricePerUnit, price)

		total := float64(piece) * price
		totalPerUnit = append(totalPerUnit, total)

		fmt.Println("Add more items?")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "yes"{
			break
		}
	}

	fmt.Println("What is your name?")
	fmt.Scanln(&cashierName)

	fmt.Println("How much discount will he get?")
	fmt.Scanln(&discount)

	invoice()
}

func invoice(){
	fmt.Printf("SEMICOLON STORES\nMAIN BRANCH\nLOCATION: 312, HERBERT MACAULAY WAY, SABO YABA LAGOS.\nTEL: 03293828343\nDate: %s\nCashier: %s\nCustomer Name: %s\n", currentDate.Format("02-01-2006"), cashierName, customerName)
	fmt.Println("========================================================")
	fmt.Println("ITEM\t\tQTY\tPRICE\t\tTOTAL(NGN)")
	fmt.Println("========================================================")
	for index := 0; index < len(itemsBought); index++ {
		fmt.Printf("%s\t\t%d\t%.2f\t\t%.2f\n", itemsBought[index], piecesBought[index], pricePerUnit[index], totalPerUnit[index])
	}
	fmt.Println("========================================================")
	
	for index := 0; index < len(totalPerUnit); index++ {
		subTotal += totalPerUnit[index]
	}
	fmt.Printf("\t\t\tSub total: %.2f\t\n", subTotal)
	
	discountPrice = (discount / 100) * subTotal
	fmt.Printf("\t\t\tDiscount: %.2f\t\n", discountPrice)

	vat = (17.50 / 100) * subTotal
	fmt.Printf("\t\t\tVAT: %.2f\t\n", vat)

	fmt.Println("========================================================")

	billTotal = subTotal - discountPrice + vat
	fmt.Printf("\t\t\tBill Total: %.2f\t\n", billTotal)
	fmt.Println("========================================================")
	fmt.Printf("THIS IS NOT A RECEIPT KINDLY PAY  %.2f\n", billTotal)
	fmt.Println("========================================================")
}