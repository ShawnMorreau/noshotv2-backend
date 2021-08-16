package noshotv2

type Room struct {
	name string
	game *Game
}

func NewRoom(name string) *Room {
	return &Room{
		name: name,
		game: NewGame(),
	}
}

func (room *Room) GetPlayersInRoom() []string {
	return room.game.PlayersArray
}

func (room *Room) JoinRoom(player Player) {
	room.game.Register <- player.(*Human)
}

func (room *Room) LeaveRoom(player Player) {
	room.game.Unregister <- player.(*Human)
}
