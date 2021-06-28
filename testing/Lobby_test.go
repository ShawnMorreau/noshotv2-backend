package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

func TestAddPlayer(t *testing.T) {
	t.Run("Add two new players to the lobby", func(t *testing.T) {
		lobby := noshotv2.NewLobby()
		lobby.AddPlayer(noshotv2.NewHuman("JimBob"))
		lobby.AddPlayer(noshotv2.NewHuman("BobJim"))

		assertLobbySize(t, lobby.Size(), 2)
	})
}
func TestRemovePlayer(t *testing.T) {
	t.Run("Remove a player from the lobby", func(t *testing.T) {
		p1 := noshotv2.NewHuman("bob")

		players := map[*noshotv2.Human]bool{
			p1:                       true,
			noshotv2.NewHuman("jim"): true,
		}
		lobby := &noshotv2.Lobby{
			Players: players,
		}
		lobby.RemovePlayer(p1)
		assertLobbySize(t, lobby.Size(), 1)
	})
}

func assertLobbySize(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d players in lobby, got %d", want, got)
	}
}
