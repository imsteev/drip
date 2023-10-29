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

func (c *Controller) GetMainPage(res *Res, req *Req) error {
	return newIndex(0, nil).Render(res)
}

func (c *Controller) GetSpace(res *Res, req *Req) error {
	spaceID, err := req.urlParamInt("spaceID")
	if err != nil {
		return fmt.Errorf("could not get param: %v", err)
	}
	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		return fmt.Errorf("could not find spaces: %v", err)
	}
	res.pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	return newIndex(spaceID, msgs).Render(res)
}

func (c *Controller) NewSpace(res *Res, req *Req) error {
	spaceID := c.SpaceGateway.Create()
	res.pushUrl(fmt.Sprintf("/spaces/%d", spaceID))
	return newIndex(spaceID, nil).Render(res)
}

func (c *Controller) CreateMessage(res *Res, req *Req) error {
	spaceID, err := req.urlParamInt("spaceID")
	if err != nil {
		return fmt.Errorf("could not get param: %v", err)
	}

	if err := c.MessageGateway.Create(spaceID, req.FormValue("text")); err != nil {
		return fmt.Errorf("could not create message: %v", err)
	}

	msgs, err := c.MessageGateway.FindBySpaceID(spaceID)
	if err != nil {
		return fmt.Errorf("error finding messages: %v", err)
	}

	return newIndex(spaceID, msgs).Render(res)
}

func (c *Controller) DeleteMessage(res *Res, req *Req) error {
	msgID, err := req.urlParamInt("messageID")
	if err != nil {
		return fmt.Errorf("could not get param: %v", err)
	}
	return c.MessageGateway.DeleteByID(msgID)
}

func newIndex(spaceID int, msgs []*models.Message) templates.Index {
	return templates.Index{
		Messages: msgs,
		SpaceID:  spaceID,
		RoomURL:  fmt.Sprintf("%s/spaces/%d", BASE_URL, spaceID),
	}
}
