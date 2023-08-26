package data

type Store struct{}

var (
	spaces = make(map[int][]string)

	MY_SPACE = 1 // simplify initial development
)

func (s Store) AddMessage(msg string, spaceID int) {
	if msg == "" {
		return
	}
	spaces[spaceID] = append(spaces[spaceID], msg)
}

func (s Store) GetMessages(spaceID int) []string {
	msgs, ok := spaces[spaceID]
	if !ok {
		return nil
	}
	return msgs
}

func (s Store) DeleteMessage(msg string, spaceID int) {
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
