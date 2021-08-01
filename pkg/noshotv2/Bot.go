package noshotv2

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
)

type Bot struct {
	ID    string
	store *PlayerCardsStore
	Game  *Game
}

func NewBot(game *Game) *Bot {
	return &Bot{
		ID:    randomdata.SillyName() + "_bot",
		store: NewCardStore(),
		Game:  game,
	}
}

func (bot *Bot) GetID() string {
	return bot.ID
}
func (bot *Bot) PlayCards(cards []string, Type string) {

}
func (bot *Bot) SetStore(store *PlayerCardsStore) {
	bot.store = store
}
func (bot *Bot) GetStore() *PlayerCardsStore {
	return bot.store
}

func (bot *Bot) Read() {
	defer func() {
		bot.Game.RemoveGenericPlayer(bot)
	}()
	for {
		if bot.Game.PlayersArray[bot.Game.PlayerAndAction.Turn] == bot.ID {
			switch bot.Game.PlayerAndAction.Action {
			case PICK_WINNER:
				rand.Seed(time.Now().UnixNano())
				winner := bot.Game.PlayersArray[rand.Intn(len(bot.Game.PlayersArray)-1)]
				bot.Game.handleWinner(fmt.Sprintf("%s%s", winner, WINNER_DELIMITER))
			case PLAY_OP:
				cards := selectRandomCard(bot.store.OP, 2)
				PlayCards(cards, OP, bot)
				bot.Game.Table.UpdateTable(bot.GetID(), cards, OP)
				mimicPayload := buildPayload(cards, OP_DELIMITER)
				bot.Game.handleCards(mimicPayload, OP_DELIMITER)
			case PLAY_NOSHOT:
				cards := selectRandomCard(bot.store.NoShot, 1)
				PlayCards(cards, NO_SHOT, bot)
				bot.Game.Table.UpdateTable(bot.GetID(), cards, NO_SHOT)
				mimicPayload := buildPayload(cards, NO_SHOT_DELIMITER)
				bot.Game.handleCards(mimicPayload, NO_SHOT_DELIMITER)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

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

func contains(cards []string, card string) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}
