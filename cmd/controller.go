package main

import (
	"drip/data"
	"drip/templates"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Controller struct {
	MessageGateway *data.MessageGateway
	SpaceGateway   *data.SpaceGateway
}

func (c *Controller) GetMainPage(w http.ResponseWriter, r *http.Request) {
	tmpl := templates.Index{}
	tmpl.MustRender(w)
}

func (c *Controller) GetSpace(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	spaceID := chi.URLParam(r, "spaceID")

	msgs, err := c.MessageGateway.GetBySpaceID(mustAtoi(spaceID))
	if err != nil {
		writeStrf(w, "%v", err)
	}

	strs := []string{}
	for _, m := range msgs {
		strs = append(strs, m.Text)
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
	tmpl := templates.Index{
		RoomURL: BASE_URL + "/spaces/" + newSpaceID,
		Space:   newSpaceID,
	}
	w.Header().Add("HX-Push-Url", "/spaces/"+newSpaceID)
	tmpl.MustRender(w)
}

func (c *Controller) CreateMessage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		writeStrf(w, "form error: %v", err)
		return
	}

	spaceID := chi.URLParam(r, "spaceID")

	c.MessageGateway.Create(mustAtoi(spaceID), r.FormValue("text"))

	msgs, err := c.MessageGateway.GetBySpaceID(mustAtoi(spaceID))
	if err != nil {
		writeStrf(w, "%v", err)
		return
	}

	strs := []string{}
	for _, m := range msgs {
		strs = append(strs, m.Text)
	}

	tmpl := templates.Index{
		Messages: strs,
		RoomURL:  BASE_URL + "/spaces/" + spaceID,
		Space:    spaceID,
	}

	tmpl.MustRender(w)
}

func (c *Controller) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	spaceID := mustAtoi(chi.URLParam(r, "spaceID"))
	if err := r.ParseForm(); err != nil {
		writeStrf(w, "form error: %v", err)
		return
	}
	c.MessageGateway.DeleteBySpaceID(spaceID)
}

func writeStrf(w io.Writer, s string, args ...any) {
	w.Write([]byte(fmt.Sprintf(s, args...)))
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
