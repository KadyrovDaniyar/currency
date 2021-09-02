package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"projects/currency/app/config"
	"projects/currency/app/db"
)

type Server struct {
	config *config.Config
	router *mux.Router
	db *db.SqlDB
	log *log.Logger
}

func Init(ctx context.Context, config *config.Config, db *db.SqlDB, log *log.Logger) *Server {
	router := mux.NewRouter()
	s :=  &Server {
		config: config,
		router: router,
		db: db,
		log: log,
	}
	s.routes()
	return s
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}

func(s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w,r)
}

