package data

import (
	"drip/data/models"
	"math/rand"
	"strconv"
)

var spaces []*models.Space

// not concurrent-safe
type SpaceGateway struct{}

func (sg *SpaceGateway) Create() *models.Space {
	// TODO: generate a random string for GUID
	id := rand.Int()
	s := &models.Space{ID: rand.Int(), GUID: strconv.Itoa(id)}
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
