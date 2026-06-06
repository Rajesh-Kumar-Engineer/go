package main

import (
	"fmt"
	"time"
)

func main() {
	var expression = 1
	// no need of break go handle internally 
	switch expression {
	case 1:
		fmt.Printf(" case 1 ")
	case 2:
		fmt.Printf(" case 2")
	default :
		fmt.Println(" default ")

	}

	// multi condition switch
	switch time.Now().Weekday(){
	case time.Sunday , time.Saturday :
		fmt.Printf("It's week end")
	default:
		fmt.Printf("It's week days")

	}
}