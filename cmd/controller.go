package main

import (
	"drip/data"
	"drip/data/models"
	"drip/templates"
	"fmt"
	"io"
	"net/http"
)

type Controller struct {
	MessageGateway *data.MessageGateway
	SpaceGateway   *data.SpaceGateway
}

func (c *Controller) GetMainPage(w http.ResponseWriter, r *http.Request) {
	newIndex(0, nil).MustRender(w)
}

func (c *Controller) GetSpace(w http.ResponseWriter, r *http.Request) {
	spaceID, err := wrapReq(r).urlParamInt("spaceID")
	if err != nil {
		writeStrf(w, "%v", err)
		return
	}

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		writeStrf(w, "%v", err)
		return
	}

	wrapRes(w).pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	newIndex(spaceID, msgs).
		MustRender(w)
}

func (c *Controller) NewSpace(w http.ResponseWriter, r *http.Request) {
	spaceID := c.SpaceGateway.Create()
	wrapRes(w).pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	newIndex(spaceID, nil).
		MustRender(w)
}

func (c *Controller) CreateMessage(w http.ResponseWriter, r *http.Request) {
	spaceID, err := wrapReq(r).urlParamInt("spaceID")
	if err != nil {
		writeStrf(w, "form error: %v", err)
		return
	}

	c.MessageGateway.Create(spaceID, r.FormValue("text"))

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		writeStrf(w, "error finding messages: %v", err)
		return
	}

	newIndex(spaceID, msgs).
		MustRender(w)
}

func (c *Controller) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	msgID, err := wrapReq(r).urlParamInt("messageID")
	if err != nil {
		writeStrf(w, "form error: %v", err)
		return
	}
	c.MessageGateway.DeleteByID(msgID)
}

func newIndex(spaceID int, msgs []*models.Message) templates.Index {
	return templates.Index{
		Messages: msgs,
		SpaceID:  spaceID,
		RoomURL:  fmt.Sprintf("%s/spaces/%d", BASE_URL, spaceID),
	}
}

func writeStrf(w io.Writer, s string, args ...any) {
	w.Write([]byte(fmt.Sprintf(s, args...)))
}
