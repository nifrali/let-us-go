package main

func main() {

	cards := createNewDeck()

	//assign the returned values of deal() to "hand" and "remainingCards" variable
	hand, remainingCards := deal(cards, 3)

	hand.printDeck()
	remainingCards.printDeck()

}