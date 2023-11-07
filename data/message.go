package data

import (
	"drip/data/models"

	"github.com/jmoiron/sqlx"
)

type MessageGateway struct {
	DB *sqlx.DB
}

func (mg *MessageGateway) Create(spaceID int, text string) error {
	mg.DB.MustExec(`INSERT INTO messages (space_id, text) VALUES (?, ?);`, spaceID, text)
	return nil
}

func (mg *MessageGateway) DeleteByID(id int) error {
	mg.DB.MustExec(`DELETE FROM messages WHERE id = ?;`, id)
	return nil
}

func (mg *MessageGateway) FindBySpaceID(spaceID int) ([]*models.Message, error) {
	var spaceMsgs []*models.Message

	rows, err := mg.DB.Query(`SELECT * FROM messages WHERE space_id = ?;`, spaceID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.SpaceID, &msg.Text); err != nil {
			return nil, err
		}
		spaceMsgs = append(spaceMsgs, &msg)
	}

	return spaceMsgs, nil
}
