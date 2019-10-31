package pets

import (
	"log"
	"os"

	"github.com/karlpokus/srv"
	"pets/internal/mongo"
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
			if err := setEnv(); err != nil {
				return err
			}
			stdout.Println("Running native. Setting env from file")
		}
		mongoHost := os.Getenv("MONGODB_HOST")
		mongoPort := os.Getenv("MONGODB_PORT")
		db, err := mongo.New("pets-service", mongoHost, mongoPort)
		if err != nil {
			return err
		}
		stdout.Println("connected to db")
		s.ExiterList = append(s.ExiterList, db)
		s.Host = os.Getenv("HTTP_HOST")
		s.Port = os.Getenv("HTTP_PORT")
		s.Router = routes(db, stdout)
		s.Logger = stdout
		return nil
	}
}
