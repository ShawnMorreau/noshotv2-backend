package noshotv2

import (
	"math/rand"
	"strings"
	"time"
)

const PLAY_NOSHOT = "Choose noShot cards"
const PICK_WINNER = "Pick a winner"
const PLAY_OP = "Choose OP cards"

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
func (g *Game) addGenericPlayer(player Player) {
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

// Cards were played
func (game *Game) handleCards(message Payload, delimiter string) {
	playerAndActionRequired := game.getPlayerAndActionRequired()
	game.Turn = playerAndActionRequired.Turn
	game.PlayerAndAction = playerAndActionRequired
	for player := range game.IPlayers {
		if !strings.Contains(player.GetID(), "_bot") {
			p := player.(*Human)
			p.Conn.WriteJSON(GameState{
				GameStarted:   true,
				Players:       game.PlayersArray,
				TurnAndAction: playerAndActionRequired,
				Judge:         game.Judge,
				Type:          5,
				Body:          "Something Happened",
				ID:            p.ID,
				Host:          game.Host.ID,
				MyOpCards:     p.store.OP,
				MyNoShotCards: p.store.NoShot,
				CardsPlayed:   game.Table.Players,
				Winner:        "",
			})
		}
	}
}
func (game *Game) handleWinner(message string) {
	winner := strings.Split(message, WINNER_DELIMITER)
	game.PlayerAndAction.Turn = 0
	for player := range game.IPlayers {
		if !strings.Contains(player.GetID(), "_bot") {
			p := player.(*Human)
			p.Conn.WriteJSON(GameState{
				GameStarted: true,
				Players:     game.PlayersArray,
				Judge:       game.Judge,
				Type:        6,
				Body:        "Picking a winner",
				CardsPlayed: game.Table.Players,
				Winner:      winner[0],
			})
		}
	}
}
func (game *Game) newRound() {
	game.Turn = game.Judge
	playerAndActionRequired := game.getPlayerAndActionRequired()
	game.Turn = playerAndActionRequired.Turn
	game.PlayerAndAction = playerAndActionRequired
	for player := range game.IPlayers {
		GrabCards(player)
		if !strings.Contains(player.GetID(), "_bot") {
			p := player.(*Human)
			p.Conn.WriteJSON(GameState{
				GameStarted:   true,
				Players:       game.PlayersArray,
				TurnAndAction: playerAndActionRequired,
				Judge:         game.Judge,
				Type:          5,
				Body:          "New Round Starting",
				ID:            p.ID,
				Host:          game.Host.ID,
				MyOpCards:     p.store.OP,
				MyNoShotCards: p.store.NoShot,
				CardsPlayed:   game.Table.Players,
				Winner:        "",
			})
		}
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
