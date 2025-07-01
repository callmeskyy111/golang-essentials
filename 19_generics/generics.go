package main

import "fmt"

func printIntSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printStrSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// printBoolSlice.. and so on.. 😪

//! Now, Generics to the rescue.. 🤩
func printSlice[T any](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}

// Narrowing them down more..
func printNarrowSlice[T int | bool](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}

// comparable interface
func comparableSlice[T comparable](items []T){
for _, item := range items {
	fmt.Println(item)

}
}

// multiple types
func multipleSlice[T comparable, V string](items []T,name V){
	for _, item := range items {
		fmt.Println(item,name)
	}
	}


//! Now, on structs
// stack DS - LIFO
type stack[T any] struct{
	elements[]T
}

func main() {
	nums:=[]int{1,2,3}
	names:=[]string{"a","b","c"}
	blaInt:=[]int{4,5,6}
	blaStr:=[]string{"GoLang","C++","JS"}
	narrowSliceBool:=[]bool{true,false,true}
	narrowSliceInt:=[]int{11,22,33}
	vals:=[]bool{true,false}
	printIntSlice(nums)
	printStrSlice(names)
	printSlice(blaInt)
	printSlice(blaStr)
	printNarrowSlice(narrowSliceBool)
	printNarrowSlice(narrowSliceInt)
	comparableSlice(vals)
	multipleSlice(vals,"Skyy")

	// Structs..
	myStackInt:= stack[int]{
		elements: []int{33,44,55},
	}

	myStackStr:= stack[string]{
		elements: []string{"Pikachu","Charizard","Raichu"},
	}

	fmt.Println(myStackInt)
	fmt.Println(myStackStr)
	
}