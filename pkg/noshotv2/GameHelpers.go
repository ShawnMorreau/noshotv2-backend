package noshotv2

import (
	"math/rand"
	"time"
)

/*
Because I made PlayersArray an array of strings(since it preserves the order everytime) I needed a way to extract the
user from a string in order to update cards
*/
func (game *Game) getClientFromName(name string) *Human {
	for client := range game.Players {
		if client.ID == name {
			return client
		}
	}
	return nil
}

// Returns a random number for the Judge when game first starts, any other judge will be determined by the player to the left
func (game *Game) getRandJudge() int {
	rand.Seed(time.Hour.Nanoseconds())
	return rand.Intn(game.Size())
}

//I disliked writing len(game.Players) everywhere so I made this
func (g *Game) Size() int {
	return len(g.Players)
}

//Add a player to our Array and Map
func (g *Game) AddPlayer(player *Human) {
	g.mu.Lock()
	if g.Size() == 0 {
		g.Host = player
	}
	g.PlayersArray = append(g.PlayersArray, player.ID)
	g.Players[player] = true
	g.mu.Unlock()
}

//Remove from Array and map
func (g *Game) RemovePlayer(player *Human) {
	g.mu.Lock()
	g.removePlayerFromArr(player)
	delete(g.Players, player)
	g.mu.Unlock()
}

//helper function for removing the player from the array
func (game *Game) removePlayerFromArr(p *Human) {
	var i int
	for idx, player := range game.PlayersArray {
		if player == p.ID {
			i = idx
		}
	}
	copy(game.PlayersArray[i:], game.PlayersArray[i+1:])
	game.PlayersArray[len(game.PlayersArray)-1] = ""
	game.PlayersArray = game.PlayersArray[:len(game.PlayersArray)-1]
}

func (game *Game) handleCards(message Payload, delimiter string) {

	playerAndActionRequired := game.getPlayerAndActionRequired()
	game.Turn = playerAndActionRequired.Turn
	for player := range game.Players {
		player.Conn.WriteJSON(GameState{
			GameStarted:   true,
			Players:       game.PlayersArray,
			TurnAndAction: playerAndActionRequired,
			Judge:         game.Judge,
			Type:          5,
			Body:          "Something Happened",
			ID:            player.ID,
			Host:          game.Host.ID,
			MyOpCards:     player.store.OP,
			MyNoShotCards: player.store.NoShot,
			CardsPlayed:   game.Table.Players,
			Winner:        "",
		})
	}

}

func (game *Game) newRound() {
	game.Turn = game.Judge
	playerAndActionRequired := game.getPlayerAndActionRequired()
	game.Turn = playerAndActionRequired.Turn

	for player := range game.Players {
		player.GrabCards()
		player.Conn.WriteJSON(GameState{
			GameStarted:   true,
			Players:       game.PlayersArray,
			TurnAndAction: playerAndActionRequired,
			Judge:         game.Judge,
			Type:          5,
			Body:          "New Round Starting",
			ID:            player.ID,
			Host:          game.Host.ID,
			MyOpCards:     player.store.OP,
			MyNoShotCards: player.store.NoShot,
			CardsPlayed:   game.Table.Players,
			Winner:        "",
		})
	}
}

func (game *Game) getPlayerToLeft(Turn int) int {
	if Turn-1 < 0 {
		return game.Size() - 1
	} else {
		return Turn - 1
	}
}

func (game *Game) getPlayerAndActionRequired() playerAndActionRequired {
	next := game.getPlayerToLeft(game.Turn)
	nextnext := game.getPlayerToLeft(next)
	if next == game.Judge {
		playersNoShotCards := game.getClientFromName(game.PlayersArray[game.Turn]).store.NoShot
		if len(playersNoShotCards) == MAX_NOSHOT_CARDS {
			return playerAndActionRequired{Turn: nextnext, Action: "Choose noShot cards", FirstToLeft: -1}
		} else {
			return playerAndActionRequired{Turn: game.Judge, Action: "Pick a winner", FirstToLeft: nextnext}
		}
	} else {
		playersOPCards := game.getClientFromName(game.PlayersArray[game.Turn]).store.OP
		if len(playersOPCards) < MAX_OP_CARDS {
			return playerAndActionRequired{Turn: next, Action: "Choose OP cards", FirstToLeft: -1}
		} else {
			return playerAndActionRequired{Turn: next, Action: "Choose noShot cards", FirstToLeft: -1}
		}
	}
}
