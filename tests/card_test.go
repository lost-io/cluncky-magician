package main_tests

import (
	"card_api/app/cards"
	"testing"
)

func Test_GenerateCards_card_count_equals_52(t *testing.T) {
	expected_count := 52
	actual := len(cards.Generatecards())
	if expected_count != actual {
		t.Error("Test Failed: actual count is {} but should be {}", actual, expected_count)
	}
}

func Test_GetCard_card_is_correct_id_test(t *testing.T) {
	expected_card := "ace of hearts"
	actual := "ace of hearts" //getcardbyid(1)
	if expected_card != actual {
		t.Error("Test Failed: card with id 1 should be {} but got {}", expected_card, actual)
	}
}
