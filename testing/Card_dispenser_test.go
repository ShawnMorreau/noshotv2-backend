package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

var fakePlayers = []noshotv2.Player{}

func TestDisperseCards(t *testing.T) {
	t.Run("Test adding cards to player with no cards", func(t *testing.T) {
		fakePlayers = append(fakePlayers, noshotv2.NewHuman("Jim"))
		assertCardsArrayLength(t)
	})
}
func assertCardsArrayLength(t testing.TB) {
	t.Helper()
	var got1, got2, want1, want2 int
	for _, player := range fakePlayers {
		player.GrabCards()
		got1 = len(player.GetStore().OP)
		got2 = len(player.GetStore().NoShot)

		want1 = 5
		want2 = 3
		if got1 != want1 {
			t.Errorf("OP Cards -> got %d, want %d", got1, want1)
		}
		if got2 != want2 {
			t.Errorf("noShot Cards -> got %d, want %d", got2, want2)
		}
	}
}