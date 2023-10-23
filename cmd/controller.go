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
	spaceID := mustAtoi(chi.URLParam(r, "spaceID"))

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		writeStrf(w, "%v", err)
	}

	tmpl := templates.Index{
		Messages: msgs,
		RoomURL:  fmt.Sprintf("%s/spaces/%d", BASE_URL, spaceID),
		SpaceID:  spaceID,
	}
	w.Header().Add("HX-Push-Url", fmt.Sprintf("/spaces/%d", spaceID))

	tmpl.MustRender(w)
}

func (c *Controller) NewSpace(w http.ResponseWriter, r *http.Request) {
	newSpaceID := rand.Int()
	tmpl := templates.Index{
		RoomURL: fmt.Sprintf("%s/spaces/%d", BASE_URL, newSpaceID),
		SpaceID: newSpaceID,
	}
	w.Header().Add("HX-Push-Url", fmt.Sprintf("/spaces/%d", newSpaceID))
	tmpl.MustRender(w)
}

func (c *Controller) CreateMessage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		writeStrf(w, "form error: %v", err)
		return
	}

	spaceID := mustAtoi(chi.URLParam(r, "spaceID"))

	c.MessageGateway.Create(spaceID, r.FormValue("text"))

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		writeStrf(w, "%v", err)
		return
	}

	tmpl := templates.Index{
		Messages: msgs,
		RoomURL:  fmt.Sprintf("%s/spaces/%d", BASE_URL, spaceID),
		SpaceID:  spaceID,
	}

	tmpl.MustRender(w)
}

func (c *Controller) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	msgID := mustAtoi(chi.URLParam(r, "messageID"))
	if err := r.ParseForm(); err != nil {
		writeStrf(w, "form error: %v", err)
		return
	}
	c.MessageGateway.DeleteByID(msgID)
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
