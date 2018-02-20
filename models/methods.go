package models

import (
	"database/sql"
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

func (db *DB) GetAllNotes(userID int64, role string) ([]note, error) {
	notes := []note{}

	var rows *sql.Rows
	var err error
	statement := "SELECT id, title, content FROM notes"
	if role == roleAdmin {
		rows, err = db.Query(statement)
	} else {
		statement += " WHERE user_id = $1"
		rows, err = db.Query(statement, userID)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot query db: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var n note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content); err != nil {
			return nil, fmt.Errorf("cannot scan response: %v", err)
		}

		notes = append(notes, n)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return notes, nil
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
