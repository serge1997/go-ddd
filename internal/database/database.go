package database

import "database/sql"

var Db *sql.DB

func New() error {
	db, err := sql.Open("", "")
	if err != nil {
		return err
	}
	Db = db
	return nil
}
