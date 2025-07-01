package main

import "fmt"

// enumerated types

// first int-types
type PikachuAttack int
const (
	ThunderBolt PikachuAttack = iota
	Thunder
	Agility
	QuickAttack
	DoubleTeam
)

// string-types
type MyFavPokemons string

const (
	Noctowl MyFavPokemons ="Noctowl"
	Umbreon ="Umbreon"
	Rapidash ="Rapidash"
)


func changePikachuAttack(attack PikachuAttack){
	fmt.Println("Pika Pika:",attack)
}

func whatsMyFavPokemon(pokemon MyFavPokemons){
	fmt.Println(pokemon," .. I choose you!")
}

func main(){
	changePikachuAttack(ThunderBolt) // Pika Pika: 0
	changePikachuAttack(Agility) // Pika Pika: 2
	whatsMyFavPokemon(Noctowl) // Noctowl  .. I choose you!
	whatsMyFavPokemon(Umbreon) // Umbreon  .. I choose you!
}