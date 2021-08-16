package noshotv2

import (
	"log"

	"github.com/gorilla/websocket"
)

type Human struct {
	ID    string
	store *PlayerCardsStore
	Conn  *websocket.Conn
	Lobby *Lobby
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

func (player *Human) Read() {
	defer func() {
		if player.Game != nil {
			player.Game.Unregister <- player
		}
		player.Conn.Close()
		player.Lobby.RemovePlayerFromLobby <- player
	}()

	for {
		messageType, p, err := player.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if player.Game != nil {
			player.Game.Broadcast <- Payload{Type: messageType, Message: string(p), User: player.ID}
		} else {
			player.Lobby.MessageBus <- GameEvent{Message: string(p), User: player.ID}
		}
	}
}
