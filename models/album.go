package models

import "fmt"

type Album struct {
	Model
	Title    string    `json:"title"`
	Slug     string    `json:"slug"`
	OwnerId  uint      `json:"owner_id"`
	Owner    *User     `json:"-"`
	Story    string    `json:"string"`
	IsActive bool      `json:"is_active"`
	Pictures []Picture `json:"-"`
}

func (a Album) String() string {
	return fmt.Sprintf("<%d: %s>", a.Id, a.Title)
}
