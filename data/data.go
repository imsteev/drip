package data

type Store struct{}

var (
	spaces = make(map[int][]string)

	MY_SPACE = 1 // simplify initial development
)

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
