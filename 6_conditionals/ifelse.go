package main

import "fmt"

func main() {
	age := 9

	if age >= 18 {
		fmt.Println("Person is an adult ✅")
	} else{
		fmt.Println("Person is NOT an adult ❌")
	}

	// Now..Else if
	pokemonType:="Electric"
	if pokemonType =="Electric"{
		fmt.Println("ThunderBolt attack!")
	}else if pokemonType=="Fire"{
		fmt.Println("Flamethrower attack!")
	}else{
		fmt.Println("WaterGun attack!")
	}

	// More
	var role = "admin"
	var hasPermissions = true
	if role == "admin" || hasPermissions {
		fmt.Println("Yes ✅")
	}

	// Declaring variables on the fly INSIDE conditionals.. ✨
	if salary:=30000;  salary >= 20000{
		fmt.Println("Decent")
	}else if salary>=15000{
		fmt.Println("Low")
	}else{
		fmt.Println("Good")
	}

	//! GOLANG DOES NOT HAVE TERNARY OPS.❌
	//! USE CONDITIONALS INSTEAD 💡

//! BETTER
// 	pokemonType:="Electric"
// 	switch pokemonType {
// case "Electric":
// 		fmt.Println("ThunderBolt attack!")
// 	case "Fire":
// 		fmt.Println("Flamethrower attack!")
// 	default:
// 		fmt.Println("WaterGun attack!")
// 	}
}