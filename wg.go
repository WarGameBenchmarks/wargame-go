package main

import (
	"fmt"
	"time"
	"math/rand"
)

// A special type to represent card Values
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

type Suit int

const (
	Clubs Suit = iota
	Hearts
	Diamonds
	Spades
)

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

type Card struct {
	s Suit
	v Value
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.v, c.s)
}

/*
	Go does not have a way to implement special
	equality and/or comparison interfaces for a struct,
	unlike Rust.

	It needs to be implemented instead.
*/
func (c Card) compare(o Card) int {
	r := 0
	switch {
		case c.v > o.v: r = 1
		case c.v < o.v: r = -1
		case c.v == o.v: r = 0
	}
	return r
}

type Deck struct {
	cards []Card
}

func (d Deck) fresh() Deck {

	values := []Value{
		Two, Three, Four, 
		Five, Six, Seven, 
		Eight, Nine, Ten, 
		Jack, Queen, King,
		Ace}
	suits := []Suit{Clubs, Hearts, Diamonds, Spades}
	
	// slice, instead of an array
	cards := make([]Card, 52)

	i := 0
	for _,suit := range suits {
		for _,value := range values {
			cards[i] = Card{suit, value}
			i++
		}
	}
	return Deck{cards}
}

func (d Deck) split() (Deck,Deck) {
	l := len(d.cards)
	h := l / 2

	part1 := d.cards[:h]
	part2 := d.cards[h:]

	return Deck{part1}, Deck{part2}
}

func (d Deck) shuffle() {
	cards := d.cards

	for i := range cards {
		j := rand.Intn(i+1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	d.cards = cards
}

/*
	Go is weird; the PRNG is not seeded automatically;
	To seed it, `init` is automatically called before main.
	Without this, the decks will always be shuffled in the same way.
*/
func init() {
	rand.Seed(time.Now().UnixNano())
}

