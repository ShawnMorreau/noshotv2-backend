package noshotv2

type Player interface {
	PlayCards(cards []string, Type string)
	GetStore() *PlayerCardsStore
	GrabCards()
	SetStore(store *PlayerCardsStore)
	Read()
}
