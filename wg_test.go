package main

import (
	"testing"
	"fmt"
)

func TestCardEquality(t *testing.T) {
	table := [][]Card{
		[]Card{Card{Hearts, Two}, Card{Hearts, Two}},
		[]Card{Card{Spades, Five}, Card{Spades, Five}},
		[]Card{Card{Hearts, Jack}, Card{Clubs, Jack}},
		[]Card{Card{Diamonds, Ace}, Card{Clubs, Ace}},
	}
	
	for _,v := range table {
		a := v[0]
		b := v[1]
		if a.compare(b) != 0 {
			t.Error("Cards are not equal")
		}		
	}
}

func TestCardNotEqual(t *testing.T) {
	table := [][]Card{
		[]Card{Card{Hearts, Six}, Card{Hearts, Two}},
		[]Card{Card{Spades, Five}, Card{Spades, Eight}},
		[]Card{Card{Clubs, Queen}, Card{Clubs, Jack}},
		[]Card{Card{Diamonds, Ace}, Card{Diamonds, King}},
	}
	
	for _,v := range table {
		a := v[0]
		b := v[1]
		if a.compare(b) == 0 {
			t.Error("Cards are equal")
		}		
	}	
}

func TestDeckFresh(t *testing.T) {
	deck := Deck{}
	deck.fresh()

	if len(deck.cards) != 52 {
		t.Error(fmt.Sprintf("Deck contains %d cards", len(deck.cards)))
	}
}

func TestDeckSplit(t *testing.T) {
	deck := Deck{}
	deck.fresh()

	p1, p2 := deck.split()

	if len(p1.cards) != 26 || len(p2.cards) != 26 {
		t.Error(fmt.Sprintf("Split decks contains %d and %d cards", len(p1.cards), len(p2.cards)))
	}
}

// Not attempting to test shuffle?