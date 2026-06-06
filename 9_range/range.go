package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4}

	// range nums --> indx , value at that idx
	for _, num := range nums {
		fmt.Println(num)
	}
	
	ageMap := map[string]int{"rajesh":24}
	
	for key , val := range ageMap{
		fmt.Println(val)
		fmt.Println(key)

	}
	// range i string
	for i , c:=range "golang"{
		fmt.Println(i)
		fmt.Println(c)
	}
}