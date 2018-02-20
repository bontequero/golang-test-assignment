package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type (
	DataLayer interface {
		GetUserInfo(login string) (*user, error)
		GetAllNotes(userID int64) []note
		GetNote(noteID int64) *note
		AddNote(userID int64) error
		DeleteNote(noteID int64) error

		// Нужно для закрытия соединения с базой данных
		Close() error
	}

	DB struct {
		*sql.DB
	}

	user struct {
		ID       int64
		Login    string
		Password string
		Role     string
	}

	note struct {
		ID      int64
		Title   string
		Content string
	}
)

func NewDB(source string) (*DB, error) {
	db, err := sql.Open("postgres", source)
	if err != nil {
		return nil, fmt.Errorf("can not open db connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can not initialize db with ping: %v", err)
	}

	return &DB{db}, nil
}
