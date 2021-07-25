package main_tests

import (
	"card_api/app/cards"
	"testing"
)

func Test_StandardDeck_GenerateCards_card_count_equals_52(t *testing.T) {
	expected_count := 52
	actual := len(cards.Generatecards())
	if expected_count != actual {
		t.Error("Test Failed: actual count is {} but should be {}", actual, expected_count)
	}
}

func Test_StandardDeck_GetCard_card_is_correct_id_test(t *testing.T) {
	var deck cards.StandardDeck
	deck.GenerateDeck(&deck.Cards)
	expected_card := 1
	actual, err := deck.GetCardById(1)
	if err != nil {
		t.Error(err)
	}
	if expected_card != actual.Id() {
		t.Error("Test Failed: card with id 1 should be {} but got {}", expected_card, actual)
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
