package data

import (
	"drip/data/models"

	"github.com/jmoiron/sqlx"
)

type SpaceGateway struct {
	DB *sqlx.DB
}

var (
	CREATE string = `INSERT INTO spaces VALUES (NULL);`
	FIND   string = `SELECT * FROM spaces WHERE id = ?;`
	DELETE string = `DELETE FROM spaces WHERE id = ?;`
)

func (sg *SpaceGateway) Create() int {
	createdID, err := sg.DB.MustExec(CREATE).LastInsertId()
	if err != nil {
		panic(err)
	}
	return int(createdID)
}

func (sg *SpaceGateway) FindByID(id int) (*models.Space, error) {
	row := sg.DB.QueryRow(FIND, id)
	var space models.Space
	if err := row.Scan(&space.ID); err != nil {
		return nil, err
	}
	return &space, nil
}

func (sg *SpaceGateway) DeleteByID(id int) {
	sg.DB.MustExec(DELETE, id)
}
