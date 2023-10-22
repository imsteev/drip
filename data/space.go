package data

import (
	"drip/data/models"
	"math/rand"
)

var spaces []*models.Space

type SpaceGateway struct{}

// not concurrent-safe
func (sg *SpaceGateway) Create() *models.Space {
	s := &models.Space{ID: rand.Int(), GUID: "asdfasdfasdfasdfasdf"}
	spaces = append(spaces, s)
	return s
}

func (sg *SpaceGateway) DeleteByID(id int) {
	updated := spaces
	for _, s := range spaces {
		copied := s
		if copied.ID != id {
			updated = append(updated, copied)
		}
	}
}

func (sg *SpaceGateway) Get(id int) *models.Space {
	for _, s := range spaces {
		if s.ID == id {
			return s
		}
	}
	return nil
}
