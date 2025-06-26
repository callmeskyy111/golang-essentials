package main

import "fmt"

//arays- numbered sequence of specific length
func main(){
var nums [4]int
var vals[4]bool
var darkTypes[3]string

//adding element
nums[0] = 11
vals[2] = true
darkTypes[0]="Umbreon"

//! zeroed vals:
// int -> 0, string -> "", bool -> false

//directly declaring, on the fly
amounts:=[3]int{100,150,140}

// 2️⃣D Arrays
twoDArray:= [2][2]int{{3,4},{5,6}}

//💡 When to use Arrays[] ❓❓❓
// - For fixed sizes 📦
// - Memory optimization 🧠
// - Constant time access (fast) ⚡

// 🟢 SLICES are more popular in GoLang


fmt.Println(vals) //[false false true false]
fmt.Println(len(nums)) //4
fmt.Println(nums[0]) //11
fmt.Println(nums) //[11 0 0 0]
fmt.Println(darkTypes) //[Umbreon  ]
fmt.Println(amounts) //[100 150 140]
fmt.Println(twoDArray) //[[3 4] [5 6]]

}