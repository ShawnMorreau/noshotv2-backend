package noshotv2

type Player interface {
	PlayCards(cards []string, Type string)
	GetStore() *PlayerCardsStore
	Read()
	SetStore(store *PlayerCardsStore)
	GetID() string
}

func PlayCards(cards []string, Type string, p Player) {
	playCards(cards, Type, p)
}
func playCards(cards []string, Type string, p Player) {
	idx := -1
	if Type == "OP" {
		for _, card := range cards {
			idx = p.GetStore().getIndexOfCard(card, p.GetStore().OP)
			copy(p.GetStore().OP[idx:], p.GetStore().OP[idx+1:])
			p.GetStore().OP[len(p.GetStore().OP)-1] = Card{}
			p.GetStore().OP = p.GetStore().OP[:len(p.GetStore().OP)-1]
		}
	} else {
		for _, card := range cards {
			idx = p.GetStore().getIndexOfCard(card, p.GetStore().NoShot)
			copy(p.GetStore().NoShot[idx:], p.GetStore().NoShot[idx+1:])
			p.GetStore().NoShot[len(p.GetStore().NoShot)-1] = Card{}
			p.GetStore().NoShot = p.GetStore().NoShot[:len(p.GetStore().NoShot)-1]
		}
	}
}
