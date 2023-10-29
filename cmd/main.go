package main

import (
	"drip/data"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var (
	PORT     = os.Getenv("PORT")
	BASE_URL = os.Getenv("BASE_URL")
)

func init() {
	if PORT == "" {
		PORT = "8080"
	}
	if BASE_URL == "" {
		BASE_URL = "http://localhost:8080"
	}
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	ctrl := Controller{
		MessageGateway: &data.MessageGateway{},
		SpaceGateway:   &data.SpaceGateway{},
	}

	r.Post("/spaces", ctrl.NewSpace)
	r.Get("/spaces/{spaceID}", ctrl.GetSpace)
	r.Post("/spaces/{spaceID}/messages", ctrl.CreateMessage)
	r.Delete("/messages/{messageID}", ctrl.DeleteMessage)
	r.Get("/", ctrl.GetMainPage)

	addr := ":" + PORT
	log.Printf("listening on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server crashed: %s", err)
	}

}
