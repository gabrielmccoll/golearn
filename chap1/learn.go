package main

import (
	"fmt"
)

func main() {

	price := 102
	fmt.Println("Price is", price, "dollars.")
	taxRate := 0.08
	tax := taxRate * float64(price)
	fmt.Println("Tax is", tax, "dollars")

	total := tax + float64(price)
	fmt.Println("Cost is", total, "Dollars")
	availableFunds := 120
	fmt.Println(availableFunds, "dollar available.")
	fmt.Println("Within budget?", float64(availableFunds) > total)
}
