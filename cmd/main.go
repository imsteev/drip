package main

import (
	"drip/data"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	PORT = "8080"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	ctrl := Controller{
		Store: new(data.Store),
	}

	r.Post("/drip", ctrl.CreateDrip)
	r.Delete("/drip", ctrl.DeleteDrip)
	r.Get("/space/{spaceID}", ctrl.GetSpace)
	r.Get("/", ctrl.GetMainPage)

	addr := ":" + PORT
	log.Printf("listening on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server crashed: %s", err)
	}
}
