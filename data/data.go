package data

type SpaceID string

var (
	spaces = make(map[SpaceID][]string)

	MY_SPACE SpaceID = "white-rabbit" // simplify initial development
)

type Store struct{}

func (s Store) AddSpace(spaceID string) {
	_, ok := spaces[SpaceID(spaceID)]
	if !ok {
		spaces[SpaceID(spaceID)] = []string{}
	}
}

func (s Store) AddMessage(msg string, spaceID SpaceID) {
	if msg == "" {
		return
	}
	spaces[spaceID] = append(spaces[spaceID], msg)
}

func (s Store) GetMessages(spaceID SpaceID) []string {
	msgs, ok := spaces[spaceID]
	if !ok {
		return nil
	}
	return msgs
}

func (s Store) DeleteMessage(msg string, spaceID SpaceID) {
	if spaces[spaceID] != nil {
		newMsgs := []string{}
		for _, m := range spaces[spaceID] {
			if m != msg {
				newMsgs = append(newMsgs, m)
			}
		}
		spaces[spaceID] = newMsgs
	}
}
