package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	pokemonType := "Fire"
	switch pokemonType {
	case "Electric":
		fmt.Println("ThunderBolt attack ⚡")
	case "Fire":
		fmt.Println("FlameThrower attack 🔥")
	default:
		fmt.Println("WaterGun attack 💦")
	}

	//multiple-condition switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's WEEKEND 🥤")
	default:
		fmt.Println("It's a WORKDAY 💼")		
	}

	//type-switch
	whatIsIt:= func (i interface{})  {
		switch t:= i.(type){
		case int:
			fmt.Println("It's an integer: ", t)
		case string:
			fmt.Println("It's a string: ", t)
	    case bool:
			fmt.Println("It's a boolean: ", t)
		default:
			fmt.Println("Other type: ",t)		
		}

	}

	whatIsIt("GoLang")
	whatIsIt(29)
	whatIsIt(true)
	whatIsIt(66.66)
}

// FlameThrower attack 🔥
// It's a WORKDAY 💼
// It's a string:  GoLang
// It's an integer:  29
// It's a boolean:  true
// Other type:  66.66

