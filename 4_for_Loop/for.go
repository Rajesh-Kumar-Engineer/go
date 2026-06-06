package main

import "fmt"

func main() {
	// there is only on function for looping in go

	// while
	i := 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	//infinite loop
	// for{
	// 	fmt.Println(i)
	// }

	// normal for loop

	for i := 0 ; i <10 ;i++ {
		fmt.Println(i)
	}

	// now we can use range 

	for j:=range 11{
		if j%4 == 0 {
			fmt.Println("Break ")
			break
		}else{
			continue
		}
		fmt.Println(j)
	}

	if rank := 1 ; rank == 1{
		fmt.Printf("Declared variable inside if construct ")
	}
}