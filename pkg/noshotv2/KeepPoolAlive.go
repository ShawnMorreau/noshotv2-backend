package noshotv2

import (
	"fmt"
	"time"
)

const KEEP_ALIVE_TYPE = 55
const HOST = 0
const SLEEP_TIME = 20

/*
	Heroku kills any websocket connections that hasn't had anything pushed through
	for 55 seconds, so I have a goroutine that sleeps for 54 seconds and then sends a
	blank piece of info over that I can just ignore by checking Type==KEEP_ALIVE_TYPE
*/
func (game *Game) StartPings() {
	for {
		time.Sleep(SLEEP_TIME * time.Second)
		game.KeepSocketAlive()
	}
}
func (game *Game) KeepSocketAlive() {
	if len(game.Players) == 0 {
		return
	} else {

		if err := game.Host.Conn.WriteJSON(GameState{
			Host: game.Host.ID,
			Type: KEEP_ALIVE_TYPE,
			Body: "...",
		}); err != nil {
			fmt.Println(err)
			return
		}
	}
}
