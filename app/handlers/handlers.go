package handlers

import (
	"card_api/cards"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type DeckHandler struct {
	Deck cards.Deck
}

func NewStandardDeckHandler() *DeckHandler {

	deck := cards.StandardDeck{}
	deck.GenerateDeck(&deck.Cards)

	return &DeckHandler{
		Deck: deck,
	}
}

func (dh *DeckHandler) HandleDeck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		dh.get(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
	}

}

func (dh *DeckHandler) randomDeck(w http.ResponseWriter, r *http.Request) {
	deck := dh.Deck.GetDeck()

	cards.ShuffleDeck(deck, 5)

	bytes, err := json.Marshal(deck)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (dh *DeckHandler) DeckOptions(w http.ResponseWriter, r *http.Request) {
	requestUrl := strings.Split(r.URL.String(), "/")

	if len(requestUrl) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if strings.ToLower(requestUrl[2]) == "random" {
		dh.randomDeck(w, r)
		return
	}

	//string conversion
	res, conversionError := strconv.ParseInt(requestUrl[2], 10, 64)
	if conversionError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(conversionError.Error()))
		return
	}
	ret := int(res)

	// get card from deck if exist
	card, err := dh.Deck.GetCardById(ret)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	bytes, err := json.Marshal(card)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

func (dh *DeckHandler) get(w http.ResponseWriter, r *http.Request) {
	cards := dh.Deck.GetDeck()
	bytes, err := json.Marshal(cards)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
