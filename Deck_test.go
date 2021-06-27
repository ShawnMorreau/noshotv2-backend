package noshotv2

import (
	"testing"
)

func TestPopulate(t *testing.T) {
	var got Deck
	err := ConvertJSONFileToStruct("test.json", &got)
	if err != nil {
		t.Fatalf("error occured in JSON %q", err)
	}
	want := 3

	if len(got.Cards) != want {
		t.Errorf("want %d, got %d", want, len(got.Cards))
	}
}
