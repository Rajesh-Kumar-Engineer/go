package main

import "fmt"

func main() {
	// slice is a dynamic array like vector in c++
	//uninitialized slice is nil
	var nums []int

	// create a slice of size 2 and cap --> capacity 5 
	var age = make([]int , 1 , 5)
	age = append(age, 1) // always add to last
	fmt.Println(len(nums))
	// copy function
	var nums2 = make([]int , len(age))
	copy(nums2 , age)
	nums2 = append(nums2 , 1000)
	fmt.Println(age)
	fmt.Println(nums2)
	fmt.Println(cap(age))

	// slice operators 
	// from [idx : idx-1]
	// slices.Equal(nums, nums2)
	fmt.Println(nums2[0:2])
}