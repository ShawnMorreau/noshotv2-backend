package noshotv2

type Player interface {
	PlayCards(cards []Card)
	GetStore() *PlayerCardsStore
	GrabCards()
	GetID() string
}
type Human struct {
	ID    string
	store *PlayerCardsStore
}

func NewHuman(name string) *Human {
	return &Human{ID: name, store: NewCardStore()}
}
func (p *Human) GetID() string {
	return p.ID
}

type PlayerCardsStore struct {
	OP     []Card
	noShot []Card
}

func NewCardStore() *PlayerCardsStore {
	return &PlayerCardsStore{
		OP:     []Card{},
		noShot: []Card{},
	}
}

func (p *Human) GetStore() *PlayerCardsStore {
	return p.store
}
func (p *Human) PlayCards(cards []Card) {
	p.store.removeCards(cards)
}

func (store *PlayerCardsStore) removeCards(cards []Card) {
	idx := -1
	if cards[0].Type == "OP" {
		for _, card := range cards {
			idx = store.getIndexOfCard(card, store.OP)
			copy(store.OP[idx:], store.OP[idx+1:])
			store.OP[len(store.OP)-1] = Card{}
			store.OP = store.OP[:len(store.OP)-1]
		}
	} else {
		for _, card := range cards {
			idx = store.getIndexOfCard(card, store.noShot)
			copy(store.noShot[idx:], store.noShot[idx+1:])
			store.noShot[len(store.noShot)-1] = Card{}
			store.noShot = store.noShot[:len(store.noShot)-1]
		}
	}

}

func (store *PlayerCardsStore) getIndexOfCard(card Card, cardStore []Card) int {
	for i, c := range cardStore {
		if card.Value == c.Value {
			return i
		}
	}
	return -1
}
