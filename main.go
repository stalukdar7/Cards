package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	// equivalent to String x = "Ace of Spades"; Java
	//card := "Ace of Spades"
	//card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}

	cards := readFromFile("my_cards")

	fmt.Println(cards.toString())
	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()
}
