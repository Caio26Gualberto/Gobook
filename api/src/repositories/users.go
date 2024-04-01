package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	statement, err := u.db.Prepare("insert into users (name, nick, email, password) values(?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}
