package main

import (
	"card_api/app/cards"
	"fmt"
)

func main() {

	deck := cards.Generatecards()
	fmt.Println("\n Before Shuffle \n")
	for i, _ := range deck {
		fmt.Println(deck[i].Name())
	}
	fmt.Println("\n After Shuffle \n")
	deck = cards.ShuffleDeck(deck, 5)
	for i, _ := range deck {
		fmt.Println(deck[i].Name())
	}
}
