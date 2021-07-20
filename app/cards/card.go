package cards

import (
	"math/rand"
)

type Suit int
type CardValue int

const (
	heart Suit = iota + 1
	diamond
	clover
	spade
)

const (
	ace CardValue = iota + 1
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
)

func (s Suit) String() string {
	return [...]string{"heart", "diamond", "clover", "spade"}[s-1]
}
func (cv CardValue) String() string {
	return [...]string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}[cv-1]
}

type (
	StandardCard struct {
		id    int
		suit  int
		value int
	}
)

type (
	Card interface {
		Id() int
		Name() string
	}
)

func (card StandardCard) Id() int {
	return card.id
}

func (card StandardCard) Name() string {
	return CardValue.String(CardValue(card.value)) + " of " + Suit.String(Suit(card.suit))
}

func Generatecards() (deck []Card) {
	var cards []Card
	for i := 1; i < 5; i++ {
		for k := 1; k < 14; k++ {
			card := StandardCard{}
			card.value = k
			card.suit = i
			card.id = i + k - 2
			cards = append(cards, card)
		}
	}
	return cards
}

func ShuffleDeck(deck []Card, iterations int) (outdeck []Card) {

	// Prevent redundant Reshuffles
	if iterations > 25 {
		iterations = 25
	}
	// no more shuffeling
	if iterations == 0 {
		return deck
	}

	sDeck := deck
	rand.Shuffle(len(sDeck), func(i, j int) { sDeck[i], sDeck[j] = sDeck[j], sDeck[i] })

	iterations--
	return ShuffleDeck(sDeck, iterations)
}
