package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	rand.Seed(time.Now().UnixNano())
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}
	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab the money and Go!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Vault cannot be opened. Heist is dead.")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Print("You were caught by the cops. Go to jail.")
		case 1:
			isHeistOn = false
			fmt.Print("Turns out the vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Print("There was an armned civilian in the bank. You should have watched them more carefully. Do ambulances save robbers? ")
		case 3:
			isHeistOn = false
			fmt.Print("Turns out the money hadn't been delivered yet. Vault is nearly empty. The money you grabbed had ink sacks that exploded... better luck next time.")
		default:
			fmt.Print("Start the getaway car! We are outta here!")
		}
	}
	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We got $", amtStolen)
	}
	fmt.Println("isHeistOn is currently:", isHeistOn)
}
