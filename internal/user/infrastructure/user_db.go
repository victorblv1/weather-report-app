package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"user/domain"
)

type UserDB struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) FetchUserByID(id string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, preferences FROM users WHERE id = ?"
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Preferences)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("error fetching user: %v", err)
	}
	return &user, nil
}

func (u *UserDB) SaveUser(user domain.User) error {
	query := "INSERT INTO users (id, name, preferences) VALUES (?, ?, ?)"
	_, err := u.db.Exec(query, user.ID, user.Name, user.Preferences)
	if err != nil {
		return fmt.Errorf("error saving user: %v", err)
	}
	return nil
}