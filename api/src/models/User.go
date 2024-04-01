package models

import "time"

type User struct {
	Id           uint64    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Nick         string    `json:"nick,omitempty"`
	Email        string    `json:"email,omitempty"`
	Password     string    `json:"password,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
}
