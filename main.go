package main

import (
	"fmt"
)

func main() {

	full_deck := Deck{}
	full_deck.Fresh()
	full_deck.Shuffle()

	p1, p2 := full_deck.Split()

	fmt.Println(len(p1.cards), "cards")
	for _,v := range p1.cards {
		fmt.Println(v)
	}
	fmt.Println(len(p2.cards), "cards")
	fmt.Println("========")
	for _,v := range p2.cards {
		fmt.Println(v)
	}

}