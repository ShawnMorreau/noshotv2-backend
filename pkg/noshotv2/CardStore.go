package noshotv2

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

func (store *PlayerCardsStore) getIndexOfCard(card string, cardStore []Card) int {
	for i, c := range cardStore {
		if card == c.Value {
			return i
		}
	}
	return -1
}
