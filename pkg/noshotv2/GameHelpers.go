package noshotv2

import (
	"math/rand"
	"time"
)

const PLAY_NOSHOT = "Choose noShot cards"
const PICK_WINNER = "Pick a winner"
const PLAY_OP = "Choose OP cards"

//GameHelpers has helper functions so that I'm not cluttering up the game file

/*
Because I made PlayersArray an array of strings(since it preserves the order) I needed a way to extract the
user from a string in order to update cards
*/
func (game *Game) getClientFromName(name string) Player {
	for client := range game.IPlayers {
		if client.GetID() == name {
			return client
		}
	}
	return nil
}

// Returns a random number for the Judge when game first starts, any other judge will be determined by the player to the left
func (game *Game) getRandJudge() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(game.Size())
}

//I disliked writing len(game.Players) everywhere so I made this
func (g *Game) Size() int {
	return len(g.IPlayers)
}

func (g *Game) AddGenericPlayer(player Player) {
	if g.Size() == 0 {
		g.Host = player.(*Human)
	}
	g.PlayersArray = append(g.PlayersArray, player.GetID())
	g.IPlayers[player] = true
}

//Remove from Array and map
func (g *Game) RemoveGenericPlayer(player Player) {
	g.removeGenericPlayerFromArr(player)
	delete(g.IPlayers, player)
}

//helper function for removing the player from the array
func (game *Game) removeGenericPlayerFromArr(p Player) {
	var i int
	for idx, player := range game.PlayersArray {
		if player == p.GetID() {
			i = idx
		}
	}
	copy(game.PlayersArray[i:], game.PlayersArray[i+1:])
	game.PlayersArray[len(game.PlayersArray)-1] = ""
	game.PlayersArray = game.PlayersArray[:len(game.PlayersArray)-1]
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
		clientNoShotCards := game.getClientFromName(game.PlayersArray[nextnext]).GetStore().NoShot
		if len(clientNoShotCards) == MAX_NOSHOT_CARDS {
			return playerAndActionRequired{Turn: nextnext, Action: PLAY_NOSHOT, FirstToLeft: -1}
		} else {
			return playerAndActionRequired{Turn: game.Judge, Action: PICK_WINNER, FirstToLeft: nextnext}
		}
	} else {
		clientOPCards := game.getClientFromName(game.PlayersArray[next]).GetStore().OP
		if len(clientOPCards) == MAX_OP_CARDS || len(clientOPCards) == 0 {
			return playerAndActionRequired{Turn: next, Action: PLAY_OP, FirstToLeft: -1}
		} else {
			return playerAndActionRequired{Turn: next, Action: PLAY_NOSHOT, FirstToLeft: -1}
		}
	}
}
func NewPayload() Payload {
	return Payload{}
}
func (game *Game) getNextUser(user string) string {
	for i, p := range game.PlayersArray {
		if p == user {
			val := game.getPlayerToLeft(i)
			if val == game.Judge {
				return game.PlayersArray[game.getPlayerToLeft(val)]
			} else {
				return game.PlayersArray[val]
			}
		}
	}
	return ""
}
