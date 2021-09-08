package main

import (
	"fmt"
	"os"
)

func main() {
	var aceValue int
	var playerName string

	fmt.Println("Welcome to BlackJack 5000!")

	fmt.Println("Could you share your name?")
	fmt.Scan(&playerName)

	fmt.Println("What's the value of the Ace card? [1 or 11]")
	fmt.Scan(&aceValue)

	if ValidateAceValue(aceValue) {
		fmt.Println(playerName + ", let the games begin!")

		// create a new game
		game := NewGame(playerName, aceValue)



	} else {
		fmt.Println("Sorry, you can't play!")
		os.Exit(1)
	}

}
