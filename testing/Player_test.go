package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

func TestPlayCards(t *testing.T) {
	hand := &noshotv2.PlayerCardsStore{
		OP: []noshotv2.Card{
			{Value: "123", Type: "OP"},
			{Value: "456", Type: "OP"},
			{Value: "789", Type: "OP"},
			{Value: "7cfqda", Type: "OP"},
			{Value: "dwawa9", Type: "OP"},
		},
		NoShot: []noshotv2.Card{
			{Value: "abc", Type: "noShot"},
			{Value: "adwahl", Type: "noShot"},
			{Value: "fwbajkdaw", Type: "noShot"},
		},
	}

	player := &noshotv2.Human{ID: "bob"}
	player.SetStore(hand)

	t.Run("Test playing two OP cards", func(t *testing.T) {
		cardsToPlay := []string{hand.OP[1].Value, hand.OP[0].Value}
		noshotv2.PlayCards(cardsToPlay, "OP", player)

		assertCardsInHand(t, len(hand.OP), 3)
	})
	t.Run("Test Playing one noShot card", func(t *testing.T) {
		cardToPlay := []string{hand.NoShot[0].Value}
		noshotv2.PlayCards(cardToPlay, "noShot", player)

		assertCardsInHand(t, len(hand.NoShot), 2)
	})
}

func assertCardsInHand(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("Wanted to see %d cards in hand but got %d", want, got)
	}
}
