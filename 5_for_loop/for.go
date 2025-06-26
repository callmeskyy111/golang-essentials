package main

import "fmt"

// for - only type of loop in GoLang
func main(){
// while loop style
i:=1
for i<= 3{
	fmt.Println(i)
	i++
}

// infinite while loop
// for{
// 	println("Hey")
// }

// classic for-loop ✨
for j:=0; j<=3; j++{
	if j==2{
		//break
		continue
	}
	fmt.Println(j)
}

// GoLang 1.22 feature - range
for k:=range 8{
	fmt.Println(k)
}
}

//O/P
// 1
// 2
// 3
// 0
// 1
// 3
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7