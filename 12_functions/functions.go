package main

//functions are 1st class citizens in GoLang

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a,b int) int{
	return  a+b
}

//multiple return types
func getLanguages()(string,string,bool){
	return "GoLang", "Javascript", true
}

//calling functions in other functions
// func processIt(fx func(a int)int){
// fx(1)
// }

func processIt() func(a int)int{
	return func (a int)int  {
		return  a
	}
}

func main() {
	sum := add(3, 5)
	sum2:= add2(10,5)
	fmt.Println(sum,sum2)
	fmt.Println(getLanguages())

	lang1, lang2, bool1 := getLanguages()
	//_ for empty vals. so that the COMPILER doesn't complain
	fmt.Println(bool1, lang2, lang1)

	// fx:=func(a int)int  {
	// 	return 2
	// }
	fn := processIt()
	fmt.Println(fn(9)) //9
}