package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Document struct {
	Name         string
	SecurityCode string
	Content      string
}

type Documents struct {
	db *sql.DB
}

const file string = "documents.db"

const drop string = `
	DROP TABLE IF EXISTS documents;
`

const create string = `
  CREATE TABLE IF NOT EXISTS documents (
	name TEXT,
	securityCode TEXT,
	content TEXT
  );
`

const seed string = `
	INSERT INTO documents VALUES("A", "A", "A Content");
	INSERT INTO documents VALUES("B", "B", "B Content");
	INSERT INTO documents VALUES("C", "C", "C Content");
	INSERT INTO documents VALUES("Default", "", "This document has an empty security code");
`

// Loads a new SQLITE3 database connection, and resets the database
// context with the seed data. This database is only intended
// to be read from by the documents service, and is intended
// to exactly mirror the contents of the seed data.
// This is the database students will execute SQL injections against.
func NewDocuments() (*Documents, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	// drop the document table if it exists, so that
	// we're sure it's a mirror of the seed data
	if _, err := db.Exec(drop); err != nil {
		return nil, err
	}
	// create database if it doesn't already exist
	if _, err := db.Exec(create); err != nil {
		return nil, err
	}
	// insert data into database
	if _, err := db.Exec(seed); err != nil {
		return nil, err
	}
	return &Documents{
		db: db,
	}, nil
}

func (d *Documents) QueryAll() ([]Document, error) {
	rows, err := d.db.Query("SELECT * FROM documents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []Document{}
	for rows.Next() {
		i := Document{}
		err = rows.Scan(&i.Name, &i.SecurityCode, &i.Content)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}

func (d *Documents) UnsafeQuery(query string) ([]Document, error) {
	// THIS IS UNSAFE, AND VULNERABLE TO SQL INJECTIONS
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []Document{}
	for rows.Next() {
		i := Document{}
		err = rows.Scan(&i.Name, &i.SecurityCode, &i.Content)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}
	return data, nil
}
