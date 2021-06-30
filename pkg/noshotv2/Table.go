package noshotv2

const MAX_OP = 2
const MAX_NOSHOT = 1

type Table struct {
	Players []PlayedInfo
}
type PlayedInfo struct {
	ID     string
	OP     []string
	NoShot []string
}

func NewTable() *Table {
	return &Table{[]PlayedInfo{}}
}
func createPlayerInfo(name string) PlayedInfo {
	return PlayedInfo{ID: name, OP: []string{}, NoShot: []string{}}
}
func (t *Table) Initialize(players []string, judge string) {
	for _, player := range players {
		if player != judge {
			t.Players = append(t.Players, createPlayerInfo(player))
		}
	}
}

func (t *Table) UpdateTable(user string, cardsToAdd []string, Type string) {
	if Type == "OP" {
		t.appendToOP(user, cardsToAdd)
		return
	}
	t.appendToNoShot(user, cardsToAdd)
}

func (t *Table) appendToOP(user string, cards []string) {
	for i, player := range t.Players {
		if player.ID == user {
			t.Players[i].OP = cards
		}
	}
}
func (t *Table) appendToNoShot(user string, cards []string) {
	for i, player := range t.Players {
		if player.ID == user {
			t.Players[i].NoShot = cards
		}
	}
}
