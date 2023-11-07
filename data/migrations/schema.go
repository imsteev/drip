package migrations

import "github.com/jmoiron/sqlx"

// Append to this list.
var MIGRATIONS []string = []string{
	`
	CREATE TABLE migrations (id INTEGER PRIMARY KEY);
	CREATE TABLE messages (id INTEGER PRIMARY KEY AUTOINCREMENT, space_id integer, text TEXT);
	CREATE TABLE spaces (id INTEGER PRIMARY KEY AUTOINCREMENT);
	PRAGMA user_version = 1;
	`,
}

func Migrate(db *sqlx.DB) error {

	var currentVersion int
	if err := db.
		QueryRow("PRAGMA user_version").
		Scan(&currentVersion); err != nil {
		return err
	}

	for _, query := range migrationsToRun(currentVersion) {
		db.MustExec(query)
	}

	return nil
}

func migrationsToRun(currentUserVersion int) []string {
	if base := currentUserVersion + 1; base < len(MIGRATIONS) {
		return MIGRATIONS[base:]
	}
	return nil
}
