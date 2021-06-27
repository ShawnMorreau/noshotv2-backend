package noshotv2

import (
	"encoding/json"
	"os"
)

type Card struct {
	Type  string
	Value string
}
type Deck struct {
	Cards []Card
}

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
