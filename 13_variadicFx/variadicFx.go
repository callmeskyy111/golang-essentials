package main

import "fmt"

//Pass 'n' no. of params. ♾️

func sum(nums ...int)int{
	tot:=0
	for _,num:=range nums{
		tot = tot+num
	}
	return  tot
}

 //any type of params
// func sum2(vals ...interface{}){
// }

func main() {
	fmt.Println(1,2,3,66,"hello") //1 2 3 66 hello
	result:=sum(3,4,10)
	fmt.Println(result) //17
	nums:=[]int{10,10,20,10} //!💡 directly, using a SLICE
	res := sum(nums...)
	fmt.Println(res) //50
}