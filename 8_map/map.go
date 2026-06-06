package main

import "fmt"

func main() {
	// create map
	m := make(map[string]string)

	m["language"] = "goLang"
	m["area"] = "backend"
	// if key does not exist in the map it return zero
	fmt.Println(m["area"])

	delete(m , "area")

	clear(m)

	// for declaration and initialization at the same time

	ageMap := map[string]int{"rajesh":24}
	fmt.Println(ageMap)

	// val , ok
	v ,ok := ageMap["rajesh"]
	if ok {
		fmt.Println("ok")
		fmt.Println(v)
	}else{
		fmt.Println("ok")
	}

	// maps.Equal()
	
}