package cards

type Suit int
type CardValue int

type Card interface {
	Id() int
	Name() string
	Suit() string
	CardValue() string
}

// Standard Card
type StandardCard struct {
	id    int    `json: "id"`
	suit  string `json: "suit"`
	value string `json: "cardvalue"`
}

// Just for Show but should be allowed when implementing the Card Interface
type TarrotCard struct {
	id     int
	suit   int
	value  int
	arcana string
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
	return card.id
}

// returns Suit of a card
func (card StandardCard) Suit() string {
	//return Suit.String(Suit(card.suit))
	return card.suit
}

// returns CardValue of a card
func (card StandardCard) CardValue() string {
	//return CardValue.String(CardValue(card.value))
	return card.value
}

// returns name of card by optaining the suit and cardvalue values
func (card StandardCard) Name() string {
	//return CardValue.String(CardValue(card.value)) + " of " + Suit.String(Suit(card.suit))
	return card.value + " of " + card.suit
}
