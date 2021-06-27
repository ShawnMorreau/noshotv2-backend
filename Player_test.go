package noshotv2

import (
	"testing"
)

func TestPlayCards(t *testing.T) {
	hand := &PlayerCardsStore{
		OP: []Card{
			{Value: "123", Type: "OP"},
			{Value: "456", Type: "OP"},
			{Value: "789", Type: "OP"},
			{Value: "7cfqda", Type: "OP"},
			{Value: "dwawa9", Type: "OP"},
		},
		noShot: []Card{
			{Value: "abc", Type: "noShot"},
			{Value: "adwahl", Type: "noShot"},
			{Value: "fwbajkdaw", Type: "noShot"},
		},
	}

	player := &Human{ID: "bob", store: hand}

	t.Run("Test playing two OP cards", func(t *testing.T) {
		cardsToPlay := []Card{hand.OP[1], hand.OP[0]}
		player.PlayCards(cardsToPlay)

		AssertCardsInHand(t, len(hand.OP), 3)
	})
	t.Run("Test Playing one noShot card", func(t *testing.T) {
		cardToPlay := []Card{hand.noShot[0]}
		player.PlayCards(cardToPlay)

		AssertCardsInHand(t, len(hand.noShot), 2)
	})
}

func AssertCardsInHand(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("Wanted to see %d cards in hand but got %d", want, got)
	}
}
