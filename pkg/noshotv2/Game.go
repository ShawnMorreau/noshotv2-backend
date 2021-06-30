package noshotv2

import (
	"strings"
	"sync"
)

// "New Round Starting"
const NO_SHOT_DELIMITER = "@#@@$@#@"
const OP_DELIMITER = "~+~...$"
const WINNER_DELIMITER = "(k(*3@#"

type Game struct {
	Host         *Human
	Players      map[*Human]bool
	Register     chan *Human
	Unregister   chan *Human
	Broadcast    chan Payload
	mu           sync.Mutex
	PlayersArray []string
	Table        *Table
	Judge        int
	Turn         int
}
type playerAndActionRequired struct {
	Turn        int
	Action      string
	FirstToLeft int
}
type GameState struct {
	GameStarted   bool
	Players       []string
	TurnAndAction playerAndActionRequired
	Judge         int
	Type          int
	Body          string
	ID            string
	Host          string
	MyOpCards     []Card
	MyNoShotCards []Card
	CardsPlayed   []PlayedInfo
	Winner        string
}

func NewGame() *Game {
	return &Game{
		Players:    make(map[*Human]bool),
		Register:   make(chan *Human),
		Unregister: make(chan *Human),
		Broadcast:  make(chan Payload),
	}
}

func (game *Game) Start() {
	game.StartPings()
	for {
		select {
		case user := <-game.Register:
			user.store = NewCardStore()
			game.AddPlayer(user)
			game.createAndSendPlayerJoinedOrLeft(2, "User has joined...", user.ID)
		case user := <-game.Unregister:
			game.RemovePlayer(user)
			game.createAndSendPlayerJoinedOrLeft(3, "User has Left...", user.ID)
		case message := <-game.Broadcast:
			switch message.Message {
			case "start game":
				InitializeDecks()
				game.Table = NewTable()
				game.Table.Initialize(game.PlayersArray, game.PlayersArray[game.Judge])
				game.Judge = game.getRandJudge()
				game.newRound()
			case "new round":
				game.Judge = game.getPlayerToLeft(game.Judge)
				game.Table = NewTable()
				game.Table.Initialize(game.PlayersArray, game.PlayersArray[game.Judge])
				game.newRound()
			case "game ended":
				for player := range game.Players {
					//gotta reset the store or it'll act weird if you end the game and try to restart it again with a full store.
					//If I want to fix this, change getPlayerActionAndTurn to check the cards played rather than the player's
					//hand
					player.store = NewCardStore()
					player.Conn.WriteJSON(GameState{Type: 99, GameStarted: false})
				}
			default:
				if strings.Contains(message.Message, OP_DELIMITER) {
					cards := strings.Split(message.Message, OP_DELIMITER)
					game.getClientFromName(game.PlayersArray[game.Turn]).PlayCards(cards, "OP")
					game.Table.UpdateTable(message.User, cards, "OP")
					game.handleCards(message, OP_DELIMITER)
				} else if strings.Contains(message.Message, NO_SHOT_DELIMITER) {
					cards := strings.Split(message.Message, NO_SHOT_DELIMITER)
					game.getClientFromName(game.PlayersArray[game.Turn]).PlayCards(cards[:1], "noShot")
					game.Table.UpdateTable(message.User, cards[:1], "noShot")
					game.handleCards(message, NO_SHOT_DELIMITER)
				} else if strings.Contains(message.Message, WINNER_DELIMITER) {
					winner := strings.Split(message.Message, WINNER_DELIMITER)
					playerAndActionRequired := game.getPlayerAndActionRequired()
					game.Turn = game.Judge
					for player := range game.Players {
						player.Conn.WriteJSON(GameState{
							GameStarted:   true,
							Players:       game.PlayersArray,
							TurnAndAction: playerAndActionRequired,
							Judge:         game.Judge,
							Type:          6,
							Body:          "Picking a winner",
							CardsPlayed:   game.Table.Players,
							Winner:        winner[0],
						})
					}
				}
			}
		}
	}

}

func (game *Game) createAndSendPlayerJoinedOrLeft(Type int, body string, who string) {
	for client := range game.Players {
		if err := client.Conn.WriteJSON(GameState{
			Players: game.PlayersArray,
			Host:    game.Host.ID,
			Type:    Type,
			Body:    body,
			ID:      who}); err != nil {
			return
		}
	}
}
