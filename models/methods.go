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

func (db *DB) GetNote(noteID int64, userID int64, role string) (*note, error) {
	result := &note{ID: noteID}
	var err error

	if role == roleAdmin {
		err = db.QueryRow("SELECT title, content FROM notes WHERE id = $1", noteID).Scan(&result.Title, &result.Content)
	} else {
		err = db.QueryRow("SELECT title, content FROM notes WHERE id = $1 AND user_id = $2", noteID, userID).Scan(&result.Title, &result.Content)
	}
	if err != nil {
		return nil, fmt.Errorf("cannot get note: %v", err)
	}

	return result, nil
}

func (db *DB) AddNote(data map[string]interface{}) error {
	_, err := db.Exec(
		"INSERT INTO notes VALUES (default, $1, $2, $3)",
		data["userID"],
		data["title"],
		data["content"],
	)
	if err != nil {
		return fmt.Errorf("cannot add note: %v", err)
	}

	return nil
}

func (db *DB) DeleteNote(noteID int64, userID int64, role string) error {
	var err error

	if role == roleAdmin {
		_, err = db.Exec("DELETE FROM notes WHERE id = $1", noteID)
	} else {
		_, err = db.Exec("DELETE FROM notes WHERE id = $1 AND user_id = $2", noteID, userID)
	}

	if err != nil {
		return fmt.Errorf("cannot delete note: %v", err)
	}

	return nil
}
