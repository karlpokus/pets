package pets

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"pets/internal/dbmock"
)

var testTable = []struct {
	name         string
	requestBody  io.Reader
	fn           http.HandlerFunc
	status       int
	responseBody []byte
}{
	{
		"stats",
		nil,
		stats(),
		200,
		[]byte(`{"name":"pets","version":"vX.Y.Z","go_version":"go1.12.4"}`),
	},
	{
		"getPets",
		nil,
		getPets(dbmock.New()),
		200,
		[]byte("rex,bixa"),
	},
	{
		"addPet",
		bytes.NewReader([]byte(`{"name":"bixa"}`)),
		addPet(dbmock.New()),
		200,
		[]byte("bixa added"),
	},
	{
		"ping",
		nil,
		ping(dbmock.New()),
		200,
		[]byte("ok"),
	},
}

func TestRoutes(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", tt.requestBody) // method and path does not matter
			w := httptest.NewRecorder()
			tt.fn(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)

			if res.StatusCode != tt.status {
				t.Errorf("expected %d, got %d", tt.status, res.StatusCode)
			}
			if !bytes.Equal(bytes.TrimSpace(body), tt.responseBody) {
				t.Errorf("expected %s, got %s", tt.responseBody, bytes.TrimSpace(body))
			}
		})
	}
}
