package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id           uint64    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("o nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("o email é obrigatório e não pode estar em branco")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("a senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
