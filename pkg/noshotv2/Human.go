package noshotv2

import (
	"log"

	"github.com/gorilla/websocket"
)

type Human struct {
	ID    string
	store *PlayerCardsStore
	Conn  *websocket.Conn
	Game  *Game
}

func NewHuman(name string) *Human {
	return &Human{ID: name, store: NewCardStore()}
}
func (p *Human) GetID() string {
	return p.ID
}
func (p *Human) SetStore(store *PlayerCardsStore) {
	p.store = store
}
func (p *Human) GetStore() *PlayerCardsStore {
	return p.store
}
func (p *Human) PlayCards(cards []string, Type string) {
	p.playCards(cards, Type)
}
func (p *Human) playCards(cards []string, Type string) {
	idx := -1
	if Type == "OP" {
		for _, card := range cards {
			idx = p.store.getIndexOfCard(card, p.store.OP)
			copy(p.store.OP[idx:], p.store.OP[idx+1:])
			p.store.OP[len(p.store.OP)-1] = Card{}
			p.store.OP = p.store.OP[:len(p.store.OP)-1]
		}
	} else {
		for _, card := range cards {
			idx = p.store.getIndexOfCard(card, p.store.NoShot)
			copy(p.store.NoShot[idx:], p.store.NoShot[idx+1:])
			p.store.NoShot[len(p.store.NoShot)-1] = Card{}
			p.store.NoShot = p.store.NoShot[:len(p.store.NoShot)-1]
		}
	}
}

func (player *Human) Read() {
	defer func() {
		player.Game.Unregister <- player
		player.Conn.Close()
	}()

	for {
		messageType, p, err := player.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		player.Game.Broadcast <- Payload{Type: messageType, Message: string(p), User: player.ID}
	}
}
