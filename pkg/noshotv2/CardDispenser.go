package noshotv2

import (
	"math/rand"
	"time"
)

const MAX_OP_CARDS = 5
const MAX_NOSHOT_CARDS = 3

func (p *Human) GrabCards() {
	p.store.OP = topUpCards(p.store.OP, "OP")
	p.store.NoShot = topUpCards(p.store.NoShot, "NOSHOT")
}

func topUpCards(cards []Card, Type string) []Card {
	switch Type {
	case "OP":
		randCards := OPDeck.GetRandomCardsFromDeck(MAX_OP_CARDS - len(cards))
		cards = append(cards, randCards...)
		return cards

	case "NOSHOT":
		for len(cards) < MAX_NOSHOT_CARDS {
			randCards := NoShotDeck.GetRandomCardsFromDeck(MAX_NOSHOT_CARDS - len(cards))
			cards = append(cards, randCards...)
			return cards
		}
		return cards
	}
	return []Card{}
}

func (deck *Deck) GetRandomCardsFromDeck(numCards int) []Card {
	rand.Seed(time.Now().UnixNano())
	var cards []Card
	if len(deck.Cards) == 0 || len(deck.Cards) < numCards {
		InitializeDecks()
	}
	for len(cards) < numCards {
		randNum := rand.Intn(len(deck.Cards))
		randCard := deck.Cards[randNum]
		cards = append(cards, randCard)
		deck.removeCardFromDeck(randNum)
	}
	return cards
}

func (deck *Deck) removeCardFromDeck(idx int) {
	copy(deck.Cards[idx:], deck.Cards[idx+1:])
	deck.Cards[len(deck.Cards)-1] = Card{}
	deck.Cards = deck.Cards[:len(deck.Cards)-1]
}
