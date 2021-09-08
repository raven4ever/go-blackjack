package main

import (
	"fmt"
	"os"
)

func main() {
	var aceValue int

	fmt.Println("Welcome to BlackJack 5000!")

	fmt.Println("What's the value of the Ace card?")
	fmt.Scan(&aceValue)

	if ValidateAceValue(aceValue) {
		fmt.Println("Let the games begin!")

		// create a new Deck
		workingDec := NewDeck(aceValue)

		// Shuffle the deck
		workingDec.Shuffle()

		fmt.Println(workingDec.String())

	} else {
		fmt.Println("Sorry, you can't play!")
		os.Exit(1)
	}

}
