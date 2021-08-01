package noshotv2

import "strings"

func (game *Game) AddBots() {

	bot := NewBot(game)
	go bot.Read()
	game.addGenericPlayer(bot)
	game.createAndSendPlayerJoinedOrLeft(2, "Bot has joined...", bot.GetID())
}
func (game *Game) RemoveBots() {
	for _, p := range game.PlayersArray {
		if strings.Contains(p, "_bot") {
			game.RemoveGenericPlayer(game.getClientFromName(p))
			game.createAndSendPlayerJoinedOrLeft(3, "Bot has left...", p)
			break
		}
	}
}
func (game *Game) RemoveAllBots() {
	for _, p := range game.PlayersArray {
		if strings.Contains(p, "_bot") {
			game.RemoveGenericPlayer(game.getClientFromName(p))
		}
	}
	game.createAndSendPlayerJoinedOrLeft(3, "Bot has left...", p)
}
