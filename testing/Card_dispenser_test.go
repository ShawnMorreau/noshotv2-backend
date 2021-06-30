package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

var fakePlayers = []noshotv2.Player{}
var fakeDeck = noshotv2.Deck{
	Cards: []noshotv2.Card{
		{Value: "abcd", Type: "OP"},
		{Value: "adwacd", Type: "OP"},
		{Value: "abfwadd", Type: "OP"},
		{Value: "abdwafd", Type: "OP"},
		{Value: "abfwaccd", Type: "OP"},
		{Value: "adwdwaacd", Type: "OP"},
		{Value: "abfcwac2awadd", Type: "OP"},
		{Value: "abdadwawafd", Type: "OP"},
		{Value: "abcd", Type: "OP"},
		{Value: "adwacd", Type: "OP"},
		{Value: "abfwadd", Type: "OP"},
		{Value: "abdwafd", Type: "OP"},
		{Value: "abcd", Type: "OP"},
		{Value: "adwacd", Type: "OP"},
		{Value: "abfwadd", Type: "OP"},
		{Value: "abdwafd", Type: "OP"},
	},
}

func TestDisperseCards(t *testing.T) {
	t.Run("Test adding cards to player with no cards", func(t *testing.T) {
		fakePlayers = append(fakePlayers, noshotv2.NewHuman("Jim"))
		assertCardsArrayLength(t)
	})
}
func assertCardsArrayLength(t testing.TB) {
	t.Helper()
	cards := fakeDeck.GetRandomCardsFromDeck(5)
	got := len(cards)
	want := 5
	if got != want {
		t.Errorf("OP Cards -> got %d, want %d", got, want)
	}
}
