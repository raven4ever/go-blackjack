package main

import (
	"fmt"
	"math/rand"
	"time"
)

// struct to model a card
type Card struct {
	value int
	suit  string
}

// struct to model a deck of cards
type Deck struct {
	cards []Card
}

// function to output a card
func (c Card) String() string {
	return fmt.Sprintf("%v of %v", c.value, c.suit)
}

// function to output a deck of cards
func (d Deck) String() string {
	var output string
	for _, card := range d.cards {
		output += fmt.Sprintf("%v\n", card)
	}
	return output
}

// function to create a new deck of cards
func NewDeck() Deck {
	var cards []Card
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, Card{value, suit})
		}
	}
	return Deck{cards}
}

// function to shuffle a deck of cards
func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	cards := d.cards

	for i := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
}
