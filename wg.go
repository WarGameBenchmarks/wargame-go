package main

import (
	"fmt"
	"time"
	"math/rand"
)

// Value represents the card value.
type Value int

const (
	Two Value = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

// String will print the value constant of a card,
// which is useful for debugging purposes.
func (v Value) String() string {
	s := ""
	switch v {
		case Two: s = "2"
		case Three: s = "3"
		case Four: s = "4"
		case Five: s = "5"
		case Six: s = "6"
		case Seven: s = "7"
		case Eight: s = "8"
		case Nine: s = "9"
		case Ten: s = "10"
		case Jack: s = "Jack"
		case Queen: s = "Queen"
		case King: s = "King"
		case Ace: s = "Ace"
	}
	return s
}

// Suit represents the card suit.
type Suit int

const (
	Clubs Suit = iota
	Hearts
	Diamonds
	Spades
)

// String will print the name of the suit constant of a card,
// which is useful for debugging.
func (u Suit) String() string {
	s := ""
	switch u {
		case Clubs: s = "Clubs"
		case Hearts: s = "Hearts"
		case Diamonds: s = "Diamonds"
		case Spades: s = "Spades"
	}
	return s
}

// Card represents a single unit in a deck,
// comprised of a Suit and Value.
type Card struct {
	s Suit
	v Value
}

// String will print the formal name of a card,
// in the format of `X of Y`. For example,
// 5 of Hearts or Ace of Spades.
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.v, c.s)
}

/*
	Go does not have a way to implement special
	equality and/or comparison interfaces for a struct,
	unlike Rust.

	It needs to be implemented instead.
*/

// Compare allows for a card to be compared against another card
// with the traditional -1,0,1 return strategy.
func (c Card) Compare(o Card) int {
	r := 0
	switch {
		case c.v > o.v: r = 1
		case c.v < o.v: r = -1
		case c.v == o.v: r = 0
	}
	return r
}

// Deck represents a collection of cards.
type Deck struct {
	cards []Card
}

// Fresh populates a deck with a `fresh` set of cards.
func (d *Deck) Fresh() {
	values := []Value{
		Two, Three, Four, 
		Five, Six, Seven, 
		Eight, Nine, Ten, 
		Jack, Queen, King,
		Ace}
	suits := []Suit{Clubs, Hearts, Diamonds, Spades}
	
	// slice, instead of an array
	d.cards = make([]Card, 52)

	i := 0
	for _,suit := range suits {
		for _,value := range values {
			d.cards[i] = Card{suit, value}
			i++
		}
	}
}

// Split attempts to evenly divide a deck in half, and returns two decks.
// The halfway point is defined with integer division,
// so a deck with 11 cards will be split [5|6]
func (d *Deck) Split() (Deck,Deck) {
	l := len(d.cards)
	h := l / 2

	part1 := d.cards[:h]
	part2 := d.cards[h:]

	return Deck{part1}, Deck{part2}
}

// Shuffle randomizes a deck- a collection of cards.
func (d *Deck) Shuffle() {
	cards := d.cards

	for i := range cards {
		j := rand.Intn(i+1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	d.cards = cards
}

func (d *Deck) GetCard() Card {
	return d.cards[0]
}

// GiveCard will give a card from this deck to another deck.
func (d *Deck) GiveCard(other *Deck) {
	// make a local copy of the deck's cards
	cards := d.cards

	// split up the slice
	// x contains the first card, cards now contains the rest
	x, cards := cards[0], cards[1:]

	// assign the rest back to this deck's cards
	d.cards = cards

	// append the other deck's cards
	other.cards = append(other.cards, x)
}

// GiveCards will give multiple cards, i.e. all of this deck's
// cards to the other deck.
func (d *Deck) GiveCards(other *Deck) {
	// shuffle will randomize the deck before its contents
	// are given to the other deck
	// TODO: this is a direct port; but feels like poor taste now
	// d.Shuffle()
	// Disabled for now

	// because `GiveCard` alters the length,
	// len(...) will recalculate incorrectly
	// so cache the size before the alterations instead.
	l := len(d.cards)

	for i := 0; i < l; i++ {
		d.GiveCard(other)
	}
}

/*
	The PRNG is not seeded automatically;
	To seed it, `init` is automatically called before main.
	Without this, the decks will always be shuffled in the same way.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

