package main

import "fmt"

func main(){
	// create a new Deck
	workingDec := NewDeck()

	// shuffle the deck
	workingDec.shuffle()

	fmt.Println(workingDec.String())
}
