package pets

import (
	"testing"
	"bytes"

	"pets/internal/dbmock"
	"github.com/karlpokus/routest"
)

func TestRoutes(t *testing.T) {
	routest.Test(t, []routest.Data{
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
	})
}
