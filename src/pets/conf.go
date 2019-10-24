package pets

import (
	"log"
	"os"

	"github.com/karlpokus/srv"
	"go.elastic.co/apm/module/apmhttprouter"
	"pets/internal/db"
)

func Logging(native bool) *log.Logger {
	var l *log.Logger
	if native {
		l = log.New(os.Stdout, "pets ", log.Ldate|log.Ltime)
	} else {
		l = log.New(os.Stdout, "", 0)
	}
	return l
}

func Conf(native bool, stdout *log.Logger) srv.ConfFunc {
	return func(s *srv.Server) error {
		if native {
			f, err := os.Open(".env")
			if err != nil {
				return err
			}
			defer f.Close()
			if err := setEnv(f); err != nil {
				return err
			}
			stdout.Println("Running native. Setting env from file")
		}
		client, err := db.New(os.Getenv("MONGODB_HOST"), os.Getenv("MONGODB_PORT"))
		if err != nil {
			return err
		}
		s.ExiterList = append(s.ExiterList, client)
		stdout.Println("connected to db")
		router := apmhttprouter.New() // wraps httprouter
		router.Handler("GET", "/api/v1/pets", logRequest(stdout, getPets(client)))
		router.Handler("POST", "/api/v1/pet", logRequest(stdout, addPet(client)))
		router.Handler("GET", "/api/v1/ping", ping(client))
		s.Router = router
		s.Host = os.Getenv("HTTP_HOST")
		s.Port = os.Getenv("HTTP_PORT")
		return nil
	}
}
