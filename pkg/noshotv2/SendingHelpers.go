package noshotv2

import (
	"log"
	"strings"
)

//If the function sends any kind of data back to the client(s). It's likely in here

func (game *Game) createAndSendPlayerJoinedOrLeft(Type int, body string, who string) {
	for client := range game.IPlayers {
		if !strings.Contains(client.GetID(), "_bot") {
			if err := client.(*Human).Conn.WriteJSON(GameState{
				Players: game.PlayersArray,
				Host:    game.Host.ID,
				Type:    Type,
				Body:    body,
				ID:      who}); err != nil {
				return
			}
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

type LobbyEvent struct {
	MyGame *Game
	ID     string
	Games  []GameInfo
}
type GameInfo struct {
	Host       string
	NumPlayers int
}

func (lobby *Lobby) sendLobbyEvent(Type int, body, who string) {
	for client := range lobby.players {
		if err := client.Conn.WriteJSON(LobbyEvent{
			// Games: lobby.buildGameInfo(),
			ID: who,
		}); err != nil {
			log.Fatalln(err)
			return
		}
	}
}

func (lobby *Lobby) buildGameInfo() []GameInfo {
	var gameInfo = []GameInfo{}
	for _, game := range lobby.games {
		gameInfo = append(gameInfo, GameInfo{Host: game.Host.GetID(), NumPlayers: len(game.PlayersArray)})
	}
	return gameInfo
}
