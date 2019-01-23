package server

import (
	"dimitrisCBR/bookie-api/v2/config"
	"dimitrisCBR/bookie-api/v2/mongodb"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router *mux.Router
	config *config.ServerConfig
}

func NewServer(u mongodb.UserService, as mongodb.AppointmentService, config config.Config) *Server {
	s := Server{
		router: mux.NewRouter(),
		config: &config.ServerConfig}

	a := authHelper{config.AuthConfig.Secret}
	NewUserRouter(u, s.getSubrouter("/user"), &a)
	NewAppointmentRouter(as, u, s.getSubrouter("/appointment"), &a)
	return &s
}

func (s *Server) Start() {
	log.Println("Listening on port " + s.config.Port)
	if err := http.ListenAndServe(s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) getSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
