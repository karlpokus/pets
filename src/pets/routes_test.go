package pets

import (
	"bytes"
	"testing"

	"github.com/karlpokus/routest"
	"pets/internal/dbmock"
)

func TestRoutes(t *testing.T) {
	routest.Test(t, []routest.Data{
		{
			"getPets",
			nil,
			getPets(dbmock.New()),
			200,
			[]byte(`[{"name":"bixa","kind":"cat"}]`),
		},
		{
			"addPet",
			bytes.NewReader([]byte(`{"name":"bixa"}`)),
			addPet(dbmock.New()),
			201,
			[]byte("bixa added"),
		},
		{
			"ping",
			nil,
			ping(dbmock.New()),
			200,
			[]byte("pong"),
		},
	})
}
