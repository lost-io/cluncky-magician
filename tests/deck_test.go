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
