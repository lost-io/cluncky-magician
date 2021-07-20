# Card API
> by lost-io


## Description:

Basic api that retuns data based on a standard 52 card deck so no jokers..


----

## Goals for the project:

1. Implement a basic api that returns standard card data
	- Data of a full deck unshuffled
	- Data of a full deck shuffled
	- Random card
 	- Specific card by card id
2. Implement swaggerui for documentation
3. Implement basic test suit for the api and automate it with github actions



## Card data structure

```` golang
{
	id: int,
	suit: int,
	value: int
}
````

### Card Functions:
 - get card id: id of a card should always be the same due to generation
 - get card name: __suit__ + __value__ etc: ace of hearts

### Deck Functions:
- get a deck of unshuffled cards etc: hearts ace to king -> then diamonds,clovers,spades
- get a deck of shuffled cards: cards have random positions in the deck

## Test Suit:

### Api Tests:



### Card Tests:
- When generating cards the expected number of cards should be 52 if it is a standard card deck
- When Shuffling a deck the deck should have different cards i different positions. check 5 postions
- Get Random card: pick two cards for a deck the cards should not be the same.







	
