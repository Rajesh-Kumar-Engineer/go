package main

import "fmt"

// function accepting many argument

func add(nums ...int)int{
	total:=0
	for _ , val := range nums{
		total +=val
	}
	return total
}

// for any type we make empty interface

func receiveAnyTyp( anytypearg ...interface{}){
	
}
func main(){
	sum := add(1,2,3,45,5)
	fmt.Println(sum)
}