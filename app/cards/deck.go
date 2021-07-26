package cards

import (
	"math/rand"
	"time"
)

type Deck interface {
	GetDeck() []Card
	GenerateDeck(cards *[]Card)
	GetCardById(id int) Card
}

type StandardDeck struct {
	Cards []Card
}

func (deck StandardDeck) GetDeck() []Card {
	return deck.Cards
}

func (deck StandardDeck) GenerateDeck(cards *[]Card) {
	for i := 1; i < 5; i++ {
		for k := 1; k < 14; k++ {
			card := StandardCard{}
			card.Value = CardValue.String(CardValue(k))
			card.Suit = Suit.String(Suit(i))
			card.ID = i + k - 2
			*cards = append(*cards, card)
		}
	}
}

func (deck StandardDeck) GetCardById(id int) (Card, error) {
	for _, card := range deck.Cards {
		if card.Id() == id {
			return card, nil
		}
	}
	return nil, &Err_CardNotFound{}
}

// A recursive function that shuffles a deck up to 25 times
func ShuffleDeck(deck []Card, iterations int) (outdeck []Card) {

	// Prevent redundant Reshuffles
	if iterations > 25 {
		iterations = 25
	}
	// no more shuffeling and return result
	if iterations == 0 {
		return deck
	}

	// shuffle deck with function from golang standard libs
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })

	iterations--
	return ShuffleDeck(deck, iterations)
}

// Get Card form a Supplied deck
func GetRandomCardFromDeck(deck []Card) (card Card) {
	rand.Seed(time.Now().UnixNano()) //changes seed everytime this function is called
	num := rand.Intn(len(deck))
	return deck[num]
}

// Get Card from a newly generated deck
func GetRandomCard() (card Card) {
	deck := Generatecards()
	return GetRandomCardFromDeck(deck)
}

// Initial CardGen Fuction
func Generatecards() (deck []Card) {
	var cards []Card
	for i := 1; i < 5; i++ {
		for k := 1; k < 14; k++ {
			card := StandardCard{}
			card.Value = CardValue.String(CardValue(k))
			card.Suit = Suit.String(Suit(i))
			card.ID = i + k - 2
			cards = append(cards, card)
		}
	}
	return cards
}
