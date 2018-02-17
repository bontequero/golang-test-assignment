package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type (
	DataLayer interface {
		GetUserInfo(login, password string) User
		GetAllNotes(userID int64) []Note
		GetNote(id int64) Note
		AddNote(userID int64) error
		DeleteNote(id int64) error
	}

	DBwrapper struct {
		*sql.DB
	}

	User struct {
		ID    int64
		Login string
		Role  string
	}

	Note struct {
		ID      int64
		Title   string
		Content string
	}
)

func NewDB(source string) (*DBwrapper, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("can not open db connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can not initialize db with ping: %v", err)
	}

	return &DBwrapper{db}, nil
}
