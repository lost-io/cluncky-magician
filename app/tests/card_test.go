package main_tests

import (
	"card_api/cards"
	"testing"
)

func Test_StandardCard_Id__returns_int(t *testing.T) {
	expected_val := 0
	cardsuit := cards.Suit.String(cards.Suit(1))
	cardvalue := cards.CardValue.String(cards.CardValue(1))
	card := cards.StandardCard{ID: 0, Suit: cardsuit, Value: cardvalue}
	if card.Id() != expected_val {
		t.Error("Card Value did not match expectedvalue")
	}
}

func Test_StandardCard_Suit_returns_string(t *testing.T) {
	expected_val := "heart"
	cardsuit := cards.Suit.String(cards.Suit(1))
	cardvalue := cards.CardValue.String(cards.CardValue(1))
	card := cards.StandardCard{ID: 0, Suit: cardsuit, Value: cardvalue}
	if card.SuitValue() != expected_val {
		t.Error("Card Value did not match expectedvalue")
	}
}

func Test_StandardCard_CardValue_returns_string(t *testing.T) {
	expected_val := "ace"
	cardsuit := cards.Suit.String(cards.Suit(1))
	cardvalue := cards.CardValue.String(cards.CardValue(1))
	card := cards.StandardCard{ID: 0, Suit: cardsuit, Value: cardvalue}
	if card.CardValue() != expected_val {
		t.Error("Card Value did not match expectedvalue")
	}
}

func Test_StandardCard_Name_returns_string(t *testing.T) {
	expected_val := "ace of heart"
	cardsuit := cards.Suit.String(cards.Suit(1))
	cardvalue := cards.CardValue.String(cards.CardValue(1))
	card := cards.StandardCard{ID: 0, Suit: cardsuit, Value: cardvalue}
	if card.Name() != expected_val {
		t.Error("Card Value did not match expectedvalue")
	}
}

func Test_GetRandomCard_Draw_Five_Not_All_Same(t *testing.T) {
	deck := cards.Generatecards()
	card_01 := cards.GetRandomCardFromDeck(deck)
	card_02 := cards.GetRandomCardFromDeck(deck)
	card_03 := cards.GetRandomCardFromDeck(deck)
	card_04 := cards.GetRandomCardFromDeck(deck)
	card_05 := cards.GetRandomCardFromDeck(deck)
	if card_01.Id() == card_02.Id() && card_01.Id() == card_03.Id() && card_01.Id() == card_04.Id() && card_01.Id() == card_05.Id() {
		t.Error("cards where all the same should have been random")
	}
}

func Test_GetRandomCard_Returns_Card(t *testing.T) {
	card := cards.GetRandomCard()
	if card == nil {
		t.Error("Random Card did not return a card")
	}
}
