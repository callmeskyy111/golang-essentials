package main

import "fmt"

func main() {
	// var name string = "golang-str"
	var name = "golang-str inferred"
	var isInteresting = true
	var age int = 29

	// shorthand syntax - directly declaration and assignment
	name1:="GoLang ShortHand"

	// indirect, late assignment
	var name2 string
	name2 = "Late assignment"

	// Now.. Floats.
	var price float32 = 50.5

	// Shorter..
	var price1 = 50.4

	// Shorthand..
	price2:=50.3

	fmt.Println(name, isInteresting, age)
	fmt.Println(name1, name2, price, price1, price2)
}