package noshotv2

import (
	"encoding/json"
	"os"
)

type Card struct {
	Type  string
	Value string `json:"value"`
}
type Deck struct {
	Cards []Card `json:"Cards"`
}

var NoShotDeck Deck
var OPDeck Deck

//Added in case I add functionality to create your own decks down the line
func ConvertJSONFileToStruct(filename string, dataStore *Deck) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(dataStore)
	if err != nil {

		return err
	}
	return err
}
func InitializeDecks() {
	ConvertJSONFileToStruct("Flaws.json", &NoShotDeck)
	ConvertJSONFileToStruct("Perks.json", &OPDeck)
}
