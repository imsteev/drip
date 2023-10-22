package data

import (
	"drip/data/models"
	"fmt"
	"math/rand"
)

var messages []*models.Message

// not concurrent-safe
type MessageGateway struct{}

func (mg *MessageGateway) Create(spaceID int, text string) (*models.Message, error) {
	m := &models.Message{ID: rand.Int(), SpaceID: spaceID, Text: text}
	messages = append(messages, m)
	return m, nil
}

func (mg *MessageGateway) DeleteByID(id int) error {
	updated := messages
	for _, m := range messages {
		if m.ID != id {
			updated = append(updated, m)
		}
	}
	messages = updated
	return nil
}

func (mg *MessageGateway) DeleteBySpaceID(spaceID int) error {
	updated := messages
	for _, m := range messages {
		if m.SpaceID != spaceID {
			updated = append(updated, m)
		}
	}
	messages = updated
	return nil
}

func (mg *MessageGateway) Get(id int) (*models.Message, error) {
	for _, m := range messages {
		if m.ID == id {
			return m, nil
		}
	}
	return nil, fmt.Errorf("no message with id %d", id)
}

func (mg *MessageGateway) GetBySpaceID(spaceID int) ([]*models.Message, error) {
	var spaceMsgs []*models.Message
	for _, m := range messages {
		if m.SpaceID == spaceID {
			spaceMsgs = append(spaceMsgs, m)
		}
	}
	return spaceMsgs, nil
}
