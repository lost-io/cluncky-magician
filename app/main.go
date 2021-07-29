package main

import (
	"card_api/app/handlers"
	"net/http"
	//"fmt"
)

func main() {
	deckHandler := handlers.NewStandardDeckHandler()
	http.HandleFunc("/deck", deckHandler.HandleDeck)
	http.HandleFunc("/deck/", deckHandler.DeckOptions)
	err := http.ListenAndServe("localhost:8080", nil) //basic error handling welp i case the server didn't start
	if err != nil {
		panic(err)
	}
}
