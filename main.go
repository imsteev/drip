package main

import (
	"drip/data"
	"drip/templates"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	s := new(data.Store)

	r.Post("/drip/text", func(w http.ResponseWriter, r *http.Request) {
		if s.GetActiveSpace() == 0 {
			s.SetNewSpace()
		}
		s.AddMessage(r.FormValue("text"), s.GetActiveSpace())
		w.Header().Add("HX-Push", fmt.Sprintf("/drip/%d", s.GetActiveSpace()))
		renderMainPage(w)
	})

	r.Get("/drip/{id}", func(w http.ResponseWriter, r *http.Request) {
		dripID := chi.URLParam(r, "id")
		fmt.Println(dripID)
		id, err := strconv.Atoi(dripID)
		if err != nil {
			respondf(w, "error: %v%v", err, id)
			return
		}
		tmpl := templates.MeTemplate{Messages: nil}

		if err := tmpl.Render(w); err != nil {
			respondf(w, "error generating template: %v", err)
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderMainPage(w)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("server crashed: %s", err)
	}
	log.Println("listening on :3000")
}

func renderMainPage(w http.ResponseWriter) {
	index := templates.IndexTemplate{
		Messages: []string{},
	}
	if err := index.Render(w); err != nil {
		respondf(w, "error generating template: %v", err)
	}
}

func respondf(w io.Writer, s string, args ...any) {
	w.Write([]byte(fmt.Sprintf(s, args...)))
}
