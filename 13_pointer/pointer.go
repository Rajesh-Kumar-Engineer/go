package main

import "fmt"

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t

}
func main() {
	// pointer pointing to the memory address
	var a int = 2
	var b int = 5
	swap(&a, &b)
	fmt.Println("a ",a)
	fmt.Println("b ",b)

}