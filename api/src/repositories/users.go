package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (u users) GetAllByNameOrNick(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := u.db.Query("select id, name, nick, email, creationTime from users where name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreationTime); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u users) GetUserById(id uint64) (models.User, error) {
	rows, err := u.db.Query("select id, name, nick, email, creationTime from users where id = ?", id)

	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.CreationTime); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (u users) Update(id uint64, user models.User) error {
	statement, err := u.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (u users) Delete(id uint64) error {
	statement, err := u.db.Prepare("delete from users where id = ?")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}
