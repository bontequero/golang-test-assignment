package models

import (
	"fmt"
)

func (db *DB) GetUserInfo(login string) (*user, error) {
	userInfo := new(user)

	err := db.QueryRow("SELECT users.id, users.login, users.password, groups.name "+
		"FROM users INNER JOIN groups ON users.group_id = groups.id "+
		"WHERE users.login = $1",
		login).Scan(&userInfo.ID, &userInfo.Login, &userInfo.Password, &userInfo.Role)
	if err != nil {
		return nil, fmt.Errorf("failed scanning: %v", err)
	}

	return userInfo, nil
}

func (db *DB) GetAllNotes(userID int64) []note {
	return nil
}

func (db *DB) GetNote(noteID int64) *note {
	return nil
}

func (db *DB) AddNote(userID int64) error {
	return nil
}

func (db *DB) DeleteNote(noteID int64) error {
	return nil
}
