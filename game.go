package main

// struct to model a game
type Game struct {
	players []Player
	deck    Deck
}

func NewGame(playerName string, aceValue int) Game {
	dealer := Player{name: "Dealer"}
	player := Player{name: playerName}

	// create a new Deck
	workingDec := NewDeck(aceValue)
	// Shuffle the deck
	workingDec.Shuffle()

	return Game{
		deck:    workingDec,
		players: []Player{dealer, player},
	}
}
