package data

import (
	"drip/data/models"

	"github.com/jmoiron/sqlx"
)

var (
	sql_create string = `INSERT INTO messages (space_id, text) VALUES (?, ?);`
	sql_delete string = `DELETE FROM messages WHERE id = ?;`
	sql_find   string = `SELECT * FROM messages WHERE space_id = ?;`
)

type MessageGateway struct {
	DB *sqlx.DB
}

func (mg *MessageGateway) Create(spaceID int, text string) error {
	mg.DB.MustExec(sql_create, spaceID, text)
	return nil
}

func (mg *MessageGateway) DeleteByID(id int) error {
	mg.DB.MustExec(sql_delete, id)
	return nil
}

func (mg *MessageGateway) FindBySpaceID(spaceID int) ([]*models.Message, error) {
	var spaceMsgs []*models.Message

	rows, err := mg.DB.Query(sql_find, spaceID)
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
