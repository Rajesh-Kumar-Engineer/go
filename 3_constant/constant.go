package main

import "fmt"

func main() {
	//constant cant not be change  | you have to assign at the time of declaration
	const name string = "Golang"

	const age = 24
	const hasPermission = true

	// constant grouping

	const (
		productName = "iphone15"
		price       = 65000
	)

	fmt.Println(productName , price)

}