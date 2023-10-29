package main

import (
	"drip/data"
	"drip/data/models"
	"drip/templates"
	"fmt"
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
	var (
		req = wrapReq(r)
		res = wrapRes(w)
	)

	spaceID, err := req.urlParamInt("spaceID")
	if err != nil {
		res.writef("could not get param: %v", err)
		return
	}

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		res.writef("could not find spaces: %v", err)
		return
	}
	res.pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	newIndex(spaceID, msgs).MustRender(w)
}

func (c *Controller) NewSpace(w http.ResponseWriter, r *http.Request) {
	var res = wrapRes(w)

	spaceID := c.SpaceGateway.Create()
	res.pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	newIndex(spaceID, nil).MustRender(w)
}

func (c *Controller) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var (
		req = wrapReq(r)
		res = wrapRes(w)
	)
	spaceID, err := req.urlParamInt("spaceID")
	if err != nil {
		res.writef("could not get param: %v", err)
		return
	}

	c.MessageGateway.Create(spaceID, r.FormValue("text"))

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		res.writef("error finding messages: %v", err)
		return
	}

	newIndex(spaceID, msgs).MustRender(w)
}

func (c *Controller) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	msgID, err := wrapReq(r).urlParamInt("messageID")
	if err != nil {
		wrapRes(w).writef("could not get param: %v", err)
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
