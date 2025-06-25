//cannot be changed/reassigned. KW - const
//can also be declared outside func main(){}, shorthand won't work tho

package main

import "fmt"
const age  = 29

func main(){
	const name string = "Skyy" //psst! can also be inferred
	//const name string = "Skyy changed" ❌

	//Grouping multiple constants..
	const (
		host = "localhost"
		port = 5000
		
	)

	fmt.Println(age)
	fmt.Println(host,port)
}