package main

import (
	"drip/data"
	"drip/templates"
	"drip/utils"
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

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// r.Post("/drip/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	s.AddMessage(r.FormValue("text"), s.GetActiveSpace())
	// 	w.Header().Add("HX-Push", fmt.Sprintf("/drip/%d", s.GetActiveSpace()))
	// 	renderMainPage(w)
	// })

	// r.Get("/drip/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	dripID := chi.URLParam(r, "id")
	// 	fmt.Println(dripID)
	// 	id, err := strconv.Atoi(dripID)
	// 	if err != nil {
	// 		respondf(w, "error: %v%v", err, id)
	// 		return
	// 	}
	// 	tmpl := templates.MeTemplate{Messages: nil}

	// 	if err := tmpl.Render(w); err != nil {
	// 		respondf(w, "error generating template: %v", err)
	// 	}
	// })

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := templates.IndexTemplate{}
		if err := tmpl.Render(w); err != nil {
			utils.WriteStrf(w, "error generating template: %v", err)
		}
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("server crashed: %s", err)
	}
}
