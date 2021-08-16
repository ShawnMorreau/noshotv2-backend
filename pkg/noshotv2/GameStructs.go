package noshotv2

type Game struct {
	Host               *Human
	IPlayers           map[Player]bool
	AddAnyTypeOfPlayer chan Player
	Players            map[*Human]bool
	Register           chan *Human
	Unregister         chan Player
	Broadcast          chan Payload
	PlayersArray       []string
	Table              *Table
	Judge              int
	Turn               int
	PlayerAndAction    playerAndActionRequired
}

type playerAndActionRequired struct {
	Turn        int
	Action      string
	FirstToLeft int
}
type Payload struct {
	Message string
	User    string
	Type    int
}
type GameState struct {
	GameStarted   bool
	Players       []string
	TurnAndAction playerAndActionRequired
	Judge         int
	Type          int
	Body          string
	ID            string
	Host          string
	MyOpCards     []Card
	MyNoShotCards []Card
	CardsPlayed   []PlayedInfo
	Winner        string
}
