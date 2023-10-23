package data

import (
	"drip/data/models"
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
	var updated []*models.Message
	for _, m := range messages {
		if m.ID != id {
			updated = append(updated, m)
		}
	}
	messages = updated
	return nil
}

func (mg *MessageGateway) FindBySpaceID(spaceID int) ([]*models.Message, error) {
	var spaceMsgs []*models.Message
	for _, m := range messages {
		if m.SpaceID == spaceID {
			spaceMsgs = append(spaceMsgs, m)
		}
	}
	return spaceMsgs, nil
}
