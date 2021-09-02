package server

import (
	httpSwagger "github.com/swaggo/http-swagger"
	_ "projects/currency/docs"
)

func(s *Server) routes() {
	s.router.HandleFunc("/currency/save/{date}", s.CurrencySave).Methods("GET")
	s.router.HandleFunc("/currency/{date}", s.CurrencyGet).Methods("GET")
	s.router.HandleFunc("/currency/{date}/{code}", s.CurrencyGetWithCode).Methods("GET")

	s.router.HandleFunc("/swagger/index.html", httpSwagger.Handler(
		httpSwagger.URL(s.config.CurrencyURL), //The url pointing to API definition
	)).Methods("GET")
}
