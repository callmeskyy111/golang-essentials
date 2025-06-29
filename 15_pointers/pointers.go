package main

import "fmt"

// variable/ds' memory-address 📍

//by value
// func changeNum(num int){
// num = 5
// fmt.Println("In changeNum fx",num)
// }

// by ref.
func changeNum(num *int){
	*num = 5
	fmt.Println("In changeNum fx",*num)
}

func main(){
	num:=1
	changeNum(&num)
	fmt.Println("After changeNum, in main fx",num)
	fmt.Println("Mem. Address: ",num)
}