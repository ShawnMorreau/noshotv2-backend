package noshotv2

const NEW_GAME_DELIMITER = "---Create a new game---"

type Lobby struct {
	games                 []*Game
	players               map[*Human]bool
	AddPlayerToLobby      chan *Human
	RemovePlayerFromLobby chan *Human
	MessageBus            chan GameEvent
}
type GameEvent struct {
	Message string
	User    string
}

func NewLobby() *Lobby {
	return &Lobby{
		players:               make(map[*Human]bool),
		AddPlayerToLobby:      make(chan *Human),
		RemovePlayerFromLobby: make(chan *Human),
		MessageBus:            make(chan GameEvent),
	}
}

func (lobby *Lobby) GetGames() []*Game {
	return lobby.games
}
func (lobby *Lobby) GetNumPlayersInLobby() int {
	return len(lobby.players)
}
func (lobby *Lobby) GetPlayerFromName(name string) *Human {
	for client := range lobby.players {
		if client.GetID() == name {
			return client
		}
	}
	return nil
}
func (lobby *Lobby) addGame(game *Game, user string) {
	lobby.games = append(lobby.games, game)
}

func (lobby *Lobby) Start() {
	go lobby.StartPings()
	for {
		select {
		case user := <-lobby.AddPlayerToLobby:
			lobby.players[user] = true
			lobby.sendLobbyEvent(1000, "Welcome to the lobby", user.ID)

		case user := <-lobby.RemovePlayerFromLobby:
			delete(lobby.players, user)

		case Event := <-lobby.MessageBus:
			if Event.Message == NEW_GAME_DELIMITER {
				game := NewGame()
				user := lobby.GetPlayerFromName(Event.User)
				user.Game = game
				lobby.addGame(game, user.GetID())
				lobby.sendLobbyEvent(1002, "Game created", user.ID)
				go game.Start()
				game.AddAnyTypeOfPlayer <- user

			}
		}
	}
}
