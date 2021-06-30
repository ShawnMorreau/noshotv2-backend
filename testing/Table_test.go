package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

var testTable = []noshotv2.PlayedInfo{
	{
		ID:     "Shawn",
		OP:     []string{},
		NoShot: []string{},
	},
	{
		ID:     "Konh",
		OP:     []string{},
		NoShot: []string{},
	},
	{
		ID:     "Jimmy",
		OP:     []string{},
		NoShot: []string{},
	},
}

func TestUpdateTable(t *testing.T) {
	table := noshotv2.Table{Players: testTable}
	testCards := []string{
		"t1", "t2"}
	table.UpdateTable("Shawn", testCards, "OP")
	got := len(testTable[0].OP)
	want := 2

	if got != want {
		t.Errorf("wanted to add %d OP cards to table for Shawn, got %d", want, got)
	}
}

func TestInitializeTable(t *testing.T) {
	table := noshotv2.NewTable()
	players := []string{"Shawn", "Kohn", "Jimmy"}

	table.Initialize(players, "judge")
	got := table
	for _, player := range got.Players {
		if !contains(t, player.ID, players) {
			t.Errorf("Table doesn't contain all of the players")
		}
	}
}

func contains(t testing.TB, player string, players []string) bool {
	t.Helper()
	for _, p := range players {
		if player == p {
			return true
		}
	}
	t.Errorf("Looking for %s but didn't find", player)
	return false
}
