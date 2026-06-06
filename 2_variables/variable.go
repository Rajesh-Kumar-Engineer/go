package main

import "fmt"

func main() {
	// in go if we declare a variable the we have to use it or delete it
	var name_of_variable string = "golang"
	
	// if we not declare type then it automatically infer the type according to the value which we are assigning to is
	var age =  12
	
	// shorthand syntax --> this cant to be used outside the function
	isAdmin := true


	fmt.Println(name_of_variable)
	fmt.Println(age)
	fmt.Println(isAdmin)

}