package main

import (
	"drip/data"
	"drip/templates"
	"drip/utils"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Controller struct {
	Store *data.Store
}

func (c *Controller) GetMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.Index{
		Messages: c.Store.GetMessages(data.MY_SPACE),
		RoomURL:  BASE_URL + "/space/" + string(data.MY_SPACE),
	}
	if err := tmpl.Render(w); err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}
}

func (c *Controller) GetSpace(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	spaceID := chi.URLParam(r, "spaceID")
	fmt.Println(spaceID, 1)
	tmpl := templates.Index{
		Messages: c.Store.GetMessages(data.SpaceID(spaceID)),
		RoomURL:  BASE_URL + "/space/" + string(data.MY_SPACE),
	}
	if err := tmpl.Render(w); err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}
}

func (c *Controller) CreateDrip(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteStrf(w, "form error: %v", err)
		return
	}
	c.Store.AddMessage(r.FormValue("text"), data.MY_SPACE)

	tmpl := templates.Index{
		Messages: c.Store.GetMessages(data.MY_SPACE),
		RoomURL:  BASE_URL + "/space/" + string(data.MY_SPACE),
	}
	if err := tmpl.Render(w); err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}
}

func (c *Controller) DeleteDrip(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteStrf(w, "form error: %v", err)
		return
	}
	text := r.FormValue("text")
	c.Store.DeleteMessage(text, data.MY_SPACE)

}
