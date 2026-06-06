package main

import (
	"fmt"
	"time"
)

// struct embedding
type customer struct {
	name string
	phone string
} 
type order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time
	customer
}

// initial setup like constructor
func newOrder(id string , amount float64 , status string)*order{
	myOrder := order{
		id : id,
		amount:amount,
		status: status,
	}
	return &myOrder
}

// attaching method to struct 
func (o *order)changeStatus(status string){
	o.status = status
}


func main() {
	// 1. Declare obj1 (all fields start at their zero-value)
	var obj1 order 

	// 2. Assign values later, one by one
	obj1.id = "1"
	obj1.amount = 45.0
	obj1.status = "pending"
	obj1.createdAt = time.Now() // Sets to current time
	fmt.Println( obj1)
	// assign by bulk
	obj1 = order{
		id : "10",
		amount: 450,
		status: "clear",
		createdAt: time.Now(),
	}
	obj2 := order{id: "10", amount: 450, status: "clear", createdAt: time.Now()}

	obj2.changeStatus("Paid")
	fmt.Println( obj1)
	fmt.Println( obj2)

	// we can also make inline struct if we wont use again ie. if we won't make it instances
	
	language := struct {
		name string
		id int
	}{"golang" ,4}
	fmt.Println(language)

	// struct embedded struct
	newOrder := order{
		id : "10",
		amount: 450,
		status: "clear",
		customer:customer {
			name : "Rajesh",
			phone : "9876543210",
		},

	}
	fmt.Println(newOrder)
}
