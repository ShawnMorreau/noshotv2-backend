package noshotv2

import (
	"strings"
)

const (
	NO_SHOT_DELIMITER    = "@#@@$@#@"
	OP_DELIMITER         = "~+~...$"
	WINNER_DELIMITER     = "(k(*3@#"
	ADD_BOT_DELIMITER    = ";'wp';"
	REMOVE_BOT_DELIMITER = ";[];"
	OP                   = "OP"
	NO_SHOT              = "noShot"
)

func NewGame() *Game {
	return &Game{
		IPlayers:           make(map[Player]bool),
		AddAnyTypeOfPlayer: make(chan Player),
		Players:            make(map[*Human]bool),
		Register:           make(chan *Human),
		Unregister:         make(chan Player),
		Broadcast:          make(chan Payload),
	}
}
func (game *Game) Start() {
	for {
		select {
		//adding players to structures that will hold Humans and Bots
		case user := <-game.AddAnyTypeOfPlayer:
			//Maybe check if store is set so it's not set again?
			user.SetStore(NewCardStore())
			game.AddGenericPlayer(user)
			game.createAndSendPlayerJoinedOrLeft(2, "User has joined...", user.GetID())
			//remove said players
		case user := <-game.Unregister:
			game.RemoveGenericPlayer(user)
			if user.GetID() == game.Host.ID {
				game.RemoveAllBots()
				game.chooseNewHost()
			}
			game.createAndSendPlayerJoinedOrLeft(3, "User has Left...", user.GetID())
		//This one is more of a generic catch, but if a message comes in from the channel..
		//do something
		case message := <-game.Broadcast:
			switch message.Message {
			//setup the game and send back the gamestate to players
			case "start game":
				InitializeDecks()
				game.Table = NewTable()
				game.Judge = game.getRandJudge()
				game.Table.Initialize(game.PlayersArray, game.PlayersArray[game.Judge])
				game.newRound()
			//reset gamestate and send back to players
			case "new round":
				for player := range game.IPlayers {
					GrabCards(player)
				}
				game.Judge = game.getPlayerToLeft(game.Judge)
				game.Table = NewTable()
				game.Table.Initialize(game.PlayersArray, game.PlayersArray[game.Judge])
				game.newRound()
			//tell client that the game is over
			case "game ended":
				for player := range game.IPlayers {
					if !strings.Contains(player.GetID(), "_bot") {
						p := player.(*Human)
						p.store = NewCardStore()
						p.Conn.WriteJSON(GameState{Type: 99, GameStarted: false})
					} else {
						player.SetStore(NewCardStore())
					}
				}
			//all defaults are for handling a played card, picking a winner, or adding bots
			default:
				if strings.Contains(message.Message, OP_DELIMITER) {
					cards := strings.Split(message.Message, OP_DELIMITER)
					PlayCards(cards, OP, game.getClientFromName(message.User))
					game.Table.UpdateTable(message.User, cards, OP)
					game.handleCards(message, OP_DELIMITER)
				} else if strings.Contains(message.Message, NO_SHOT_DELIMITER) {
					cards := strings.Split(message.Message, NO_SHOT_DELIMITER)
					PlayCards(cards[:1], NO_SHOT, game.getClientFromName(message.User))
					nextUser := game.getNextUser(message.User)
					game.Table.UpdateTable(nextUser, cards[:1], NO_SHOT)
					game.handleCards(message, NO_SHOT_DELIMITER)

				} else if strings.Contains(message.Message, ADD_BOT_DELIMITER) {
					game.AddBots()
				} else if strings.Contains(message.Message, REMOVE_BOT_DELIMITER) {
					game.RemoveBots()
				} else if strings.Contains(message.Message, WINNER_DELIMITER) {
					game.handleWinner(message.Message)
				}
			}
		}
	}
}
