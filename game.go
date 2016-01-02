package main

import (
	"math/rand"
)

func Game(generator *rand.Rand) {

	deck := NewDeckWithGenerator([]Card{}, generator)
	deck.Fresh()
	deck.Shuffle()

	player1, player2 := deck.Split()

	turns := 0

	// an empty deck
	winner := NewDeckWithGenerator([]Card{}, generator)

	base: for len(player1.cards) > 0 && len(player2.cards) > 0 {
		turns += 1

		c1, c2 := player1.GetCard(), player2.GetCard()

		player1.GiveCard(winner)
		player2.GiveCard(winner)

		logger.Printf("%s vs %s\n", c1, c2)

		if c1.Compare(c2) == 0 {
			wars := 0

			// not quite a do-while?
			for c1.Compare(c2) == 0 {
				if len(player1.cards) < 4 || len(player2.cards) < 4 {
					break base
				}

				wars += 1

				logger.Printf("War %d\n", wars)

				for i := 0; i < 3; i++ {
					player1.GiveCard(winner)
					player2.GiveCard(winner)
				}

				c1, c2 = player1.GetCard(), player2.GetCard()

				player1.GiveCard(winner)
				player2.GiveCard(winner)

				logger.Printf("%s vs %s\n", c1, c2)

				if c1.Compare(c2) > 0 {
					winner.GiveCards(player1)
					logger.Printf("W: P1 - %d cards\n", len(player1.cards))
				} else if c1.Compare(c2) < 0 {
					winner.GiveCards(player2)
					logger.Printf("W: P2 - %d cards\n", len(player2.cards))
				} else {
					// another war
				}

			}

		} else if c1.Compare(c2) > 0 {
			winner.Shuffle()
			winner.GiveCards(player1)
			logger.Printf("P1 - %d cards\n", len(player1.cards))
		} else if c1.Compare(c2) < 0 {
			winner.Shuffle()
			winner.GiveCards(player2)
			logger.Printf("P2 - %d cards\n", len(player2.cards))
		}

	}

	logger.Printf("P1 (%d) vs P2 (%d) in %d turns\n", len(player1.cards), len(player2.cards), turns)

}