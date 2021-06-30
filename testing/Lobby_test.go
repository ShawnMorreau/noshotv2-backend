package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

func TestAddPlayer(t *testing.T) {
	t.Run("Add two new players to the game", func(t *testing.T) {
		game := noshotv2.NewGame()
		game.AddPlayer(noshotv2.NewHuman("JimBob"))
		game.AddPlayer(noshotv2.NewHuman("BobJim"))

		assertGameSize(t, game.Size(), 2)
	})
}
func TestRemovePlayer(t *testing.T) {
	t.Run("Remove a player from the game", func(t *testing.T) {
		p1 := noshotv2.NewHuman("bob")

		players := map[*noshotv2.Human]bool{
			p1:                       true,
			noshotv2.NewHuman("jim"): true,
		}
		game := &noshotv2.Game{
			Players: players,
		}
		game.RemovePlayer(p1)
		assertGameSize(t, game.Size(), 1)
	})
}

func assertGameSize(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d players in game, got %d", want, got)
	}
}
