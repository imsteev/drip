package data

import "math/rand"

type Store struct{}

var activeSpace int
var spaces = make(map[int][]string)

func (s Store) AddMessage(msg string, spaceID int) {
	spaces[spaceID] = append(spaces[spaceID], msg)
}

func (s Store) GetMessages(spaceID int) []string {
	msgs, ok := spaces[spaceID]
	if !ok {
		return nil
	}
	return msgs
}

func (s Store) GetActiveSpace() int {
	return activeSpace
}

func (s Store) SetNewSpace() int {
	space := newSpace()
	activeSpace = space
	return space
}

func newSpace() int {
	for {
		n := rand.Int()
		if _, found := spaces[n]; found {
			continue
		}
		spaces[n] = []string{}
		return n
	}
}
