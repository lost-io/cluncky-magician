package cards

type Suit int
type CardValue int

type Card interface {
	Id() int
	Name() string
	SuitValue() string
	CardValue() string
}

// Standard Card
type StandardCard struct {
	ID    int    `json: "id"`
	Suit  string `json: "suit"`
	Value string `json: "cardvalue"`
}

// Just for Show but should be allowed when implementing the Card Interface
type TarrotCard struct {
	Id     int
	Suit   int
	Value  int
	Arcana string
}

//StandardDeck Suit Constants
const (
	heart Suit = iota + 1
	diamond
	clover
	spade
)

//StandardDeck CardValue Constants
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

// binder for the Suit type to const
func (s Suit) String() string {
	return [...]string{"heart", "diamond", "clover", "spade"}[s-1]
}

// binder for the CardValue to const
func (cv CardValue) String() string {
	return [...]string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}[cv-1]
}

// return id of card
func (card StandardCard) Id() int {
	return card.ID
}

// returns Suit of a card
func (card StandardCard) SuitValue() string {
	return card.Suit
}

// returns CardValue of a card
func (card StandardCard) CardValue() string {
	return card.Value
}

// returns name of card by optaining the suit and cardvalue values
func (card StandardCard) Name() string {
	return card.Value + " of " + card.Suit
}
