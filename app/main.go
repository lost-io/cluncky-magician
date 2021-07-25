package main

import (
	"net/http"
	//"fmt"
)

func deckHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	/*
		deck := cards.Generatecards()
		fmt.Println("\n Before Shuffle \n")
		for i, _ := range deck {
			fmt.Println(deck[i].Name())
		}
		fmt.Println("\n After Shuffle \n")
		deck = cards.ShuffleDeck(deck, 5)
		for i, _ := range deck {
			fmt.Println(deck[i].Name())
		}*/
	http.HandleFunc("/deck", deckHandler)
	err := http.ListenAndServe("localhost:8080", nil) //basic error handling welp i case the server didn't start
	if err != nil {
		panic(err)
	}
}
