package main

import (
	"fmt"
	"maps"
)

// maps -> hash_tables, objects, dicts.

func main(){
	m:=make(map[string]string)
	m1:=make(map[string]int)

	// setting elements
	m["name"] = "Raichu"
	m["type"] = "Electric ⚡"
	m["attack"] = "ThunderBolt"
	m1["age"] = 29

	// map, without "make" kw
	m2:=map[string]int{"currYear":2025,"birthYear":1995}

	//💡 If elements known beforehand: "make", otherwise NO

	// Chcking whether element exists or not!
	v,ok:=m2["currYear"]
	fmt.Println(v) //2025
	if ok{
		fmt.Println("Property exists ✅.. All OK!")
	}else{
		fmt.Println("Not OK! 🔴",v) //0
	}

	// get an element
	fmt.Println(m["name"])
	fmt.Println(m["type"])
	fmt.Println(m["nonExistingKey"]) //Empty Val.
	fmt.Println(m1["age"])
	fmt.Println(m2)

	// length
	fmt.Println(len(m))

	// deleting
	fmt.Println(m)
	delete(m,"type")
	fmt.Println(m)

	// clearing the map
	clear(m)
	fmt.Println(m)

	//Checking equality
	m3:=map[string]int{"squirtle":7}
	m4:=map[string]int{"squirtle":7}

	fmt.Println(maps.Equal(m3,m4))

}
