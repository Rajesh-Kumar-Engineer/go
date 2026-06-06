package main

import "fmt"

func main() {
	var nums [6]int // initialized to zero by default -->zeroed value
	length := len(nums)
	fmt.Println(length)

	arr := [4]int{1,2,3,4}


	// 2d array

	studentAge := [2][2] int{{2,2},{3,4}}


	
	fmt.Println(studentAge) 
	fmt.Println(arr)
}