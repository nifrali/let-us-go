//"Receiver" design pattern
package main

import (
	"fmt"
)

//create a new type of "deck" as slice of string
type deck []string

//method that create a new deck of cards
func createNewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spade", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//replace indexes(i, j) with
	//"_" since they're not used
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//method looping through the deck and printing values
func (d deck) printDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//method returning 2 decks of cards (multiple return values)
func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}