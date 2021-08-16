package noshotv2

import (
	"fmt"
	"time"
)

const (
	KEEP_ALIVE_TYPE = 55
	SLEEP_TIME      = 50
)

/*
	Heroku kills any websocket connections that hasn't had anything pushed through
	for 55 seconds, so I have a goroutine that sleeps for 54 seconds and then sends a
	blank piece of info over that I can just ignore by checking Type==KEEP_ALIVE_TYPE
*/
func (lobby *Lobby) StartPings() {
	for {
		time.Sleep(SLEEP_TIME * time.Second)
		lobby.KeepSocketAlive()
	}
}
func (lobby *Lobby) KeepSocketAlive() {
	if len(lobby.players) == 0 {
		return
	} else {

		if err := lobby.GetPlayerFromLobby().Conn.WriteJSON(GameState{
			Type: KEEP_ALIVE_TYPE,
			Body: "...",
		}); err != nil {
			fmt.Println(err)
			return
		}
	}
}
func (lobby *Lobby) GetPlayerFromLobby() *Human {
	var key *Human
	for k := range lobby.players {
		key = k
		break
	}
	return key
}
