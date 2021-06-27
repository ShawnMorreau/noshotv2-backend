package noshotv2

import (
	"fmt"
	"sync"
)

type Lobby struct {
	Host    Human
	Players map[*Human]bool
	mu      sync.Mutex
}

func NewLobby() *Lobby {
	return &Lobby{
		Players: make(map[*Human]bool),
	}
}

func (l *Lobby) AddPlayer(player *Human) {
	l.mu.Lock()
	fmt.Println(player)
	l.Players[player] = true
	l.mu.Unlock()
}

func (l *Lobby) Size() int {
	return len(l.Players)
}
