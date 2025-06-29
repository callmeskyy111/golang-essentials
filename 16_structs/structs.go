// custom DS
// blueprint
package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	phone string
}

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time // A Time represents an instant in time with nanosecond precision.
	customer // struct-embedding
}

// constructor fx (newStructName: convention)
func newOrder(id string, amount float32,status string) *order{
	// some initial setup goes here...
		order1 := order{
			id: id,
			amount: amount,
			status: status,
		}
		return &order1
	}

// method to change status
// o - first letter of the struct's name (convention)
func(o *order) changeStatus(status string){
o.status=status
}

//! 🔴🔴🔴 To modify - use *ptr, otherwise not needed
// ZeroValues work here (string->" ", int/float->0, bool->false)

// same, to change amount
func(o order) getAmount() float32{
	return o.amount
}

func main(){
    // COMPOSITION
	language := struct {
		name string
		isGood bool
	}{
		"GoLang", true,
	}

	fmt.Println(language) // {GoLang true}

	order1:=newOrder("1",150.99,"received")
	order2 := order{
		id: "2",
		amount: 440.55,
		status: "delivered",
		createdAt: time.Now(),
	}

	// customer1:=customer{
	// 	name: "Skyy",
	// 	phone: "123456",
	// }


	// struct-embedding
	order3 :=order{
		id: "3",
		amount: 50.66,
		status: "received",
		//inline embedding
		customer: customer{
			name: "Skyy",
			phone: "123456",
		},
	}

	order3.customer.name = "robin" //changing

	fmt.Println(order3)
	fmt.Println(order3.customer)

	order1.createdAt = time.Now()
	order1.amount = 44.66
	order2.changeStatus("changedStatus")

	fmt.Println(order1.amount)
	fmt.Println(order1.getAmount())


	fmt.Println(order1)
	fmt.Println(order2)
}

