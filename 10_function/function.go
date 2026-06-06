package main

import "fmt"

// func add(a int, b int) {
// 	// take something return nothing
// 	fmt.Println(a , b)
// }
func sub()int {
	// take something return nothing
	a:= 4
	b := 5

	return b-a
}
func add(a int, b int)int {
	// take something return something
	return a+b
}

// if last argument have type and before that not type then all will be the type of last one
func multiplication(a , b int)int {
	// take something return something
	return a*b
}


// multi value return function
func getLanguage()(string , string ,string , int){
	return "golang", "c++" , "c" ,1
} 


// function accepting func as an argument
// (functionNamewhich pas tp function , func()return type of funvtion) return type of accepting func{}
func processIt(fn func(a int)int){

	fn(1);
}
// function  returning function
func process()func(a int) int{
	return func(a int)int{
		return a
	}
}
func main() {
	sum := add(2,4)
	subtract := sub()
	fmt.Println(sum)
	fmt.Println(subtract)
	fmt.Println(getLanguage())
	// we can extract from multivalued return function\
	lang1 , lang2 ,lang3 ,val := getLanguage()
	fmt.Println(lang1 , lang2 ,lang3 ,val)
	// function returning function
	function := process()
	value := function(2)
	fmt.Println(value)
}