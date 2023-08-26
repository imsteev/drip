package main

import (
	"drip/data"
	"drip/handlers"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	ctrl := handlers.Controller{
		Store: new(data.Store),
	}

	r.Post("/drip", ctrl.CreateDrip)
	r.Delete("/drip", ctrl.DeleteDrip)
	r.Get("/space/{spaceID}", ctrl.GetSpace)
	r.Get("/", ctrl.GetMainPage)

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("server crashed: %s", err)
	}
}
