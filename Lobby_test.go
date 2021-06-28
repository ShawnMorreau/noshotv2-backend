package noshotv2

import "testing"

func TestAddPlayer(t *testing.T) {
	t.Run("Add two new players to the lobby", func(t *testing.T) {
		lobby := NewLobby()
		lobby.AddPlayer(NewHuman("JimBob"))
		lobby.AddPlayer(NewHuman("BobJim"))

		assertLobbySize(t, lobby.Size(), 2)
	})
}
func TestRemovePlayer(t *testing.T) {
	t.Run("Remove a player from the lobby", func(t *testing.T) {
		p1 := NewHuman("bob")

		players := map[*Human]bool{
			p1:              true,
			NewHuman("jim"): true,
		}
		lobby := &Lobby{
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
