# Card API
> by lost-io


## Description:

Basic api that retuns data based on a standard 52 card deck so no jokers..

mainly used with curl command

etc:
* curl localhost:8080/deck gives a deck
* curl localhost:8080/deck/1 gives a card from the deck where the id matches
* curl localhost:8080/deck/random  gives a shuffled deck
----

## Make Commands:

1. runserver
	. starts an instace of the main.go which runs the api on localhost:8080

2. runtests
	. runs unit tests

3. runtestsv
	. runs unit tests with the --verbose flag

4. build
	. builds the docker image

5. compose-build
	. builds the docker image and pusheds to docker registry


## Goals for the project:

1. Implement a basic api that returns standard card data
	- [x] Data of a full deck unshuffled
	- [x] Data of a full deck shuffled
	- [] Random card
 	- [x] Specific card by card id
2. Implement swaggerui for documentation
3. Implement basic test suit for the api and automate it with github actions



## Card data structure

```` golang
{
	id: int,
	suit: string,
	value: string
}
````

### Card Functions:
 - get card id: id of a card should always be the same due to generation
 - get card name: __suit__ + __value__ etc: ace of hearts

### Deck Functions:
- [x] get a deck of unshuffled cards etc: hearts ace to king -> then diamonds,clovers,spades
- [x] get a deck of shuffled cards: cards have random positions in the deck

## Test Suit:

### Api Tests:
- [x] GetDeck
- [x] ShuffledDeck
- [x] CardByIdFromDeck


### Card Tests:
- [x] When generating cards the expected number of cards should be 52 if it is a standard card deck
- [x] When Shuffling a deck the deck should have different cards i different positions. check 5 postions
- [x] Get Random card: pick two cards for a deck the cards should not be the same.







	
