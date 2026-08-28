package cache

import (
	"context"
	"database/sql"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDatabase()  error {
	cachePath, err := getCachePath()
	dbPath := filepath.Join(cachePath, "tunectl.sqlite")
	if err != nil {
		return err
	}
	
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS artists (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			name TEXT NOT NULL, 
			artistmbid TEXT NOT NULL UNIQUE
		)`,
	)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS albums (
         id INTEGER PRIMARY KEY AUTOINCREMENT,
         artistid INTEGER NOT NULL,
         name TEXT NOT NULL,
         CONSTRAINT fk_artistid
            FOREIGN KEY (artistid) REFERENCES artists(id)
		);`,
	)
	if err != nil {
		return err
	}
	
	_, err = db.ExecContext(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS songs (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			albumid INTEGER NOT NULL,
			title TEXT NOT NULL, 
			mbid TEXT NOT NULL UNIQUE,
			CONSTRAINT fk_albumid FOREIGN KEY (id) REFERENCES albums(id)
		)`,
	)
	if err != nil {
		return err
	}
	
	return nil
}