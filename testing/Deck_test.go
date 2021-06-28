package noshotv2

import (
	"testing"

	"github.com/shawnmorreau/noshotv2-backend/pkg/noshotv2"
)

func TestPopulate(t *testing.T) {
	var got noshotv2.Deck
	err := noshotv2.ConvertJSONFileToStruct("test.json", &got)
	if err != nil {
		t.Fatalf("error occured in JSON %q", err)
	}
	want := 3

	if len(got.Cards) != want {
		t.Errorf("want %d, got %d", want, len(got.Cards))
	}
}
