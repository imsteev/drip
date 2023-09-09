package main

import (
	"drip/data"
	"drip/templates"
	"drip/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Controller struct {
	Store *data.Store
}

func (c *Controller) GetMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.Index{}
	tmpl.MustRender(w)
}

func (c *Controller) GetSpace(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	spaceID := chi.URLParam(r, "spaceID")
	msgs, err := c.Store.FindMessages(data.SpaceID(spaceID))
	if err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}

	strs := []string{}
	for _, m := range msgs {
		strs = append(strs, m.Message)
	}
	tmpl := templates.Index{
		Messages: strs,
		RoomURL:  BASE_URL + "/spaces/" + spaceID,
		Space:    spaceID,
	}
	w.Header().Add("HX-Push-Url", "/spaces/"+spaceID)

	tmpl.MustRender(w)
}

func (c *Controller) NewSpace(w http.ResponseWriter, r *http.Request) {
	newSpaceID := strconv.Itoa(rand.Int())
	c.Store.AddSpace(newSpaceID)
	msgs, err := c.Store.FindMessages(data.SpaceID(newSpaceID))
	if err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}
	strs := []string{}
	for _, m := range msgs {
		strs = append(strs, m.Message)
	}
	tmpl := templates.Index{
		Messages: strs,
		RoomURL:  BASE_URL + "/spaces/" + newSpaceID,
		Space:    newSpaceID,
	}
	w.Header().Add("HX-Push-Url", "/spaces/"+newSpaceID)

	tmpl.MustRender(w)
}

func (c *Controller) CreateDrip(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteStrf(w, "form error: %v", err)
		return
	}

	space := chi.URLParam(r, "spaceID")
	if space == "" {
		space = strconv.Itoa(rand.Int())
	}

	c.Store.AddMessage(r.FormValue("text"), data.SpaceID(space))

	msgs, err := c.Store.FindMessages(data.SpaceID(space))
	if err != nil {
		utils.WriteStrf(w, "error generating template: %v", err)
	}

	strs := []string{}
	for _, m := range msgs {
		strs = append(strs, m.Message)
	}
	tmpl := templates.Index{
		Messages: strs,
		RoomURL:  BASE_URL + "/spaces/" + space,
		Space:    space,
	}

	tmpl.MustRender(w)
}

func (c *Controller) DeleteDrip(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.WriteStrf(w, "form error: %v", err)
		return
	}
	space := chi.URLParam(r, "spaceID")

	text := r.FormValue("text")
	fmt.Println(space, text)

}
