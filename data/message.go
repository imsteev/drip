package data

import (
	"drip/data/models"
	"math/rand"
)

var messages []*models.Message

type MessageGateway struct{}

// not concurrent-safe
func (mg *MessageGateway) Create(spaceID int, text string) *models.Message {
	m := &models.Message{ID: rand.Int(), SpaceID: spaceID, Text: text}
	messages = append(messages, m)
	return m
}

func (mg *MessageGateway) DeleteByID(id int) {
	updated := messages
	for _, m := range messages {
		if m.ID != id {
			updated = append(updated, m)
		}
	}
	messages = updated
}

func (mg *MessageGateway) DeleteBySpaceID(spaceID int) {
	updated := messages
	for _, m := range messages {
		if m.SpaceID != spaceID {
			updated = append(updated, m)
		}
	}
	messages = updated
}

func (mg *MessageGateway) Get(id int) *models.Message {
	for _, m := range messages {
		if m.ID == id {
			return m
		}
	}
	return nil
}

func (mg *MessageGateway) GetBySpaceID(spaceID int) []*models.Message {
	var spaceMsgs []*models.Message
	for _, m := range messages {
		if m.SpaceID == spaceID {
			spaceMsgs = append(spaceMsgs, m)
		}
	}
	return spaceMsgs
}
