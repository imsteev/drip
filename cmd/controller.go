package main

import (
	"drip/data"
	"drip/data/models"
	"drip/templates"
	"fmt"
)

type Controller struct {
	MessageGateway *data.MessageGateway
	SpaceGateway   *data.SpaceGateway
}

func (c *Controller) GetMainPage(res *Res, req *Req) {
	newIndex(0, nil).MustRender(res)
}

func (c *Controller) GetSpace(res *Res, req *Req) {
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
	newIndex(spaceID, msgs).MustRender(res)
}

func (c *Controller) NewSpace(res *Res, req *Req) {
	spaceID := c.SpaceGateway.Create()
	res.pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	newIndex(spaceID, nil).MustRender(res)
}

func (c *Controller) CreateMessage(res *Res, req *Req) {
	spaceID, err := req.urlParamInt("spaceID")
	if err != nil {
		res.writef("could not get param: %v", err)
		return
	}

	c.MessageGateway.Create(spaceID, req.FormValue("text"))

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		res.writef("error finding messages: %v", err)
		return
	}

	newIndex(spaceID, msgs).MustRender(res)
}

func (c *Controller) DeleteMessage(res *Res, req *Req) {
	msgID, err := req.urlParamInt("messageID")
	if err != nil {
		res.writef("could not get param: %v", err)
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
