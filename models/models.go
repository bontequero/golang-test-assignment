package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type (
	DataLayer interface {
		GetUserInfo(string) (*user, error)
		GetAllNotes(int64, string) ([]note, error)
		GetNote(int64, int64, string) (*note, error)
		AddNote(map[string]interface{}) error
		DeleteNote(int64, int64, string) error

		// Метод нужен для закрытия соединения с базой данных из функции main
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

const (
	roleAdmin = "admin"
	roleUser  = "user"
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
