package models

import "time"

type UserDTO struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"nome"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"criado_em"`
	Followers []UserDTO `json:"followers"`
	Following []UserDTO `json:"following"`
	Posts     []PostDTO `json:"posts"`
}
