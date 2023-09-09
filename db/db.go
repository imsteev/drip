package db

import (
	"database/sql"
	"errors"
)

type SpaceID string

type Store struct {
	DB *sql.DB
}

type Message struct {
	ID      int
	SpaceID int
	Message string
}

func (s Store) AddMessage(msg string, spaceID SpaceID) error {
	if msg == "" {
		return errors.New("message must be non-empty")
	}

	_, err := s.DB.Exec(`INSERT INTO messages (message, spaceID) VALUES (?, ?)`, msg, spaceID)
	if err != nil {
		return err
	}

	return nil
}

func (s Store) DeleteMessage(msg string, spaceID SpaceID) error {
	stmt, err := s.DB.Prepare(`DELETE FROM messages WHERE spaceID = ? AND message = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(spaceID, msg); err != nil {
		return err
	}

	return nil
}

func (s *Store) FindMessages(spaceID SpaceID) ([]*Message, error) {
	rows, err := s.DB.Query(`SELECT * FROM messages WHERE spaceID = $1`, spaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var msgs []*Message
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.Message, &msg.SpaceID); err != nil {
			return nil, err
		}
		msgs = append(msgs, &msg)
	}

	return msgs, nil
}
