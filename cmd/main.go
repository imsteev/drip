package main

import (
	"drip/data"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var s *data.Store

func init() {
	s = new(data.Store)
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/drip", CreateDrip)
	r.Delete("/drip", DeleteDrip)
	r.Get("/", GetMainPage)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("server crashed: %s", err)
	}
}
