package noshotv2

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (game *Game) AddBots() {
	bot := NewBot(game)
	go bot.Read()
	game.AddGenericPlayer(bot)
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
	game.createAndSendPlayerJoinedOrLeft(3, "Bot has left...", "someBot")
}

//handleCards looks for a Payload, so we mimic what a payload from the client would look like
func buildPayload(cards []string, delimiter string) Payload {
	payload := NewPayload()
	if delimiter == OP_DELIMITER {
		payload.Message = fmt.Sprintf("%s%s%s", cards[0], delimiter, cards[1])
	} else {
		payload.Message = fmt.Sprintf("%s%s", cards[0], delimiter)
	}
	return payload
}

func selectRandomCard(cardArr []Card, num int) []string {
	rand.Seed(time.Now().UnixNano())
	var cards []string

	for len(cards) < num {
		randInt := rand.Intn(len(cardArr))
		randCard := cardArr[randInt].Value
		if !contains(cards, randCard) {
			cards = append(cards, randCard)
		}
	}
	return cards
}
