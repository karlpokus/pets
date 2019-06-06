package pets

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	*http.Server
	Port string
}

var Version = "vX.Y.Z" // injected at build time

var Stdout = log.New(os.Stdout, "", 0)
var Stderr = log.New(os.Stderr, "", 0)

func (s *Server) Start() error {
	Stdout.Println(fmt.Sprintf("pets %s listening on port %s", Version, s.Port))
	return s.ListenAndServe()
}

func New(port string) *Server {
	router := httprouter.New()
	router.Handler("GET", "/api/v1/pets", logRequest(getPets()))
	router.Handler("GET", "/api/v1/version", logRequest(getVersion()))

	return &Server{
		Server: &http.Server{
			Addr:              fmt.Sprintf(":%s", port),
			Handler:           router,
			ReadTimeout:       10 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			MaxHeaderBytes:    1 << 20, // 1 MB
		},
		Port: port,
	}
}
