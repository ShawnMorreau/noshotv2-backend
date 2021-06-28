package noshotv2

type Player interface {
	PlayCards(cards []Card)
	GetStore() *PlayerCardsStore
	GrabCards()
	GetID() string
	SetStore(store *PlayerCardsStore)
}
type Human struct {
	ID    string
	store *PlayerCardsStore
}

func (p *Human) SetStore(store *PlayerCardsStore) {
	p.store = store
}
func NewHuman(name string) *Human {
	return &Human{ID: name, store: NewCardStore()}
}
func (p *Human) GetID() string {
	return p.ID
}

type PlayerCardsStore struct {
	OP     []Card
	NoShot []Card
}

func NewCardStore() *PlayerCardsStore {
	return &PlayerCardsStore{
		OP:     []Card{},
		NoShot: []Card{},
	}
}

func (p *Human) GetStore() *PlayerCardsStore {
	return p.store
}
func (p *Human) PlayCards(cards []Card) {
	p.store.removeCards(cards)

}

/*
	DRY here. Although, I think the only workaround would be using the reflection
	package and at that point, it's easier and probably therefore I think
	it's fine to have a little bit of repetition.
*/
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
			idx = store.getIndexOfCard(card, store.NoShot)
			copy(store.NoShot[idx:], store.NoShot[idx+1:])
			store.NoShot[len(store.NoShot)-1] = Card{}
			store.NoShot = store.NoShot[:len(store.NoShot)-1]
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
