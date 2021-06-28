package 

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHome(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	Server(res, req)

	got := res.Body.String()
	want := "Hello there"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
