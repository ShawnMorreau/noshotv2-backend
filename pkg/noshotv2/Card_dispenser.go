package noshotv2

const MAX_OP_CARDS = 5
const MAX_NOSHOT_CARDS = 3

func (p *Human) GrabCards() {
	p.store.OP = topUpCards(p.store.OP, "OP")
	p.store.NoShot = topUpCards(p.store.NoShot, "NOSHOT")
}

func topUpCards(cards []Card, Type string) []Card {
	switch Type {
	case "OP":
		for len(cards) < MAX_OP_CARDS {
			cards = append(cards, getRandomOPCard())
		}
		return cards

	case "NOSHOT":
		for len(cards) < MAX_NOSHOT_CARDS {
			cards = append(cards, getRandomNoShotCard())
		}
		return cards
	}
	return []Card{}
}

func getRandomOPCard() Card {
	return Card{Value: "abcd", Type: "OP"}
}

func getRandomNoShotCard() Card {
	return Card{Value: "abcd", Type: "noShot"}
}
