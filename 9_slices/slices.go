package main

import (
	"fmt"
	"slices"
)

// Dynamic Arrays
// Most used construct in GoLang
// +useful methods

func main(){
	// uninitialized slice - nil
	var nums[]int

	//initial size
	var pokemons =  make([]string,0,5)

	//adding
	pokemons = append(pokemons, "Umbreon 🖤","Quagsire 💦")
	pokemons = append(pokemons, "Zapdos ⚡")
	pokemons = append(pokemons, "Alakazam 🔮")
	nums = append(nums, 3)

	//adding using indices
	pokemons[3] = "Charizard 🔥"
	
	//copying fx.
	var pokemons2 = make([]string, len(pokemons))
	pokemons2 = append(pokemons2, "Bulbasaur 🌿")

	copy(pokemons2, pokemons)
	fmt.Println(pokemons, pokemons2)

	//slice operator
	var fireTypes = []string{"Charmander","Magmar","Growlith","Moltres","Rapidash"}
	fmt.Println(fireTypes[0:1])
	fmt.Println(fireTypes[1:2])
	fmt.Println(fireTypes[1:])

	// slices pkg
	var vals = []int{2,4}
	var vals2 = []int{2,4}

	fmt.Println(slices.Equal(vals,vals2))

	// 2D slices
	var  nums2D = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums2D)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(pokemons)
	fmt.Println(cap(pokemons)) //max no. of element/capacity.
	fmt.Println(pokemons==nil)
	
}

//O/P:

// [Umbreon 🖤 Quagsire 💦 Zapdos ⚡ Charizard 🔥] [Umbreon 🖤 Quagsire 💦 Zapdos ⚡ Charizard 🔥 Bulbasaur 🌿]
// [Charmander]
// [Magmar]
// [Magmar Growlith Moltres Rapidash]
// true
// [[1 2 3] [4 5 6]]
// [3]
// 1
// [Umbreon 🖤 Quagsire 💦 Zapdos ⚡ Charizard 🔥]
// 5
// false
