package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//seeding the random function
	rand.Seed(time.Now().UnixNano())

	//tells if the heist is on
	isHeistOn := true

	numberGuards := rand.Intn(100)

	//if the number of guards ar less or equal to 20 we can pass them
	if numberGuards <= 25 {
		fmt.Println("Fihom, we can do it guys")
	} else {
		fmt.Println("Mafihach, we can't do it, to much guards:", numberGuards, "guards")
		isHeistOn = false
	}

	//the resistance of the vault
	vaultResistance := rand.Intn(100)

	//if the heist is still on and the vault resistance is weak
	if isHeistOn && vaultResistance <= 55 {
		fmt.Println("Grab and Go !")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Abort! The vault can't be opened, we'll bring better tools next time !")
	}

	if isHeistOn {
		//chances to leave safely
		switch leftSafely := rand.Intn(5); leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("We are stuck in here, heist failed")
		case 1:
			isHeistOn = false
			fmt.Println("The alarm brought the police too early, heist failed")
		case 2:
			isHeistOn = false
			fmt.Println("The cameras cought our faces, heist failed")
		case 3:
			isHeistOn = false
			fmt.Println("The bags are too heavy, heist failed")
		default:
			fmt.Println("Good job, let's get the hell out of here")
		}
	}

	if isHeistOn {
		//how much was in the vault
		amountStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("We've got $%v", amountStolen)
	}

}
