package migrations

var SCHEMA string = `
	CREATE TABLE messages (id INTEGER PRIMARY KEY AUTOINCREMENT, space_id integer, text TEXT);
	CREATE TABLE spaces (id INTEGER PRIMARY KEY AUTOINCREMENT);
`
