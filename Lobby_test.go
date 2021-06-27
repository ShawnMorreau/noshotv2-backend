package noshotv2

import "testing"

func TestJoinLobby(t *testing.T) {
	lobby := NewLobby()
	lobby.AddPlayer(NewHuman("JimBob"))
	lobby.AddPlayer(NewHuman("BobJim"))

	got := lobby.Size()
	want := 2

	if got != want {
		t.Errorf("Wanted %d players, got %d", want, got)
	}
}
