package main

import (
	"drip/data"
	"drip/data/migrations"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"

	// driver for sqlx
	_ "github.com/mattn/go-sqlite3"
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

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("could not start database: %v", err)
	}
	defer db.Close()

	// FLAW: at this point
	db.MustExec(migrations.SCHEMA)

	ctrl := Controller{
		MessageGateway: &data.MessageGateway{DB: db},
		SpaceGateway:   &data.SpaceGateway{DB: db},
	}

	r.Post("/spaces", wrapped(ctrl.NewSpace))
	r.Get("/spaces/{spaceID}", wrapped(ctrl.GetSpace))
	r.Post("/spaces/{spaceID}/messages", wrapped(ctrl.CreateMessage))
	r.Delete("/messages/{messageID}", wrapped(ctrl.DeleteMessage))
	r.Get("/", wrapped(ctrl.GetMainPage))

	addr := ":" + PORT
	log.Printf("listening on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server crashed: %s", err)
	}

}

type myHandler func(res *Res, req *Req) error

func wrapped(h myHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(wrapRes(w), wrapReq(r)); err != nil {
			// TODO: how to present this without navigating user to a new page
			w.Write([]byte(err.Error()))
		}
	}
}
