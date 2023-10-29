package data

import (
	"drip/data/models"

	"github.com/jmoiron/sqlx"
)

var spaces []*models.Space

// not concurrent-safe
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

func (sg *SpaceGateway) FindByID(id int) *models.Space {
	row := sg.DB.QueryRow(FIND, id)
	var space *models.Space
	if err := row.Scan(&space); err != nil {
		panic(err)
	}
	return space
}

func (sg *SpaceGateway) DeleteByID(id int) {
	sg.DB.MustExec(DELETE, id)
}
