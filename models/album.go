package models

import (
	"fmt"
	"strings"
)

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

func (a *Album) Valid() error {
	if len(strings.TrimSpace(a.Title)) == 0 {
		return fmt.Errorf("title cannot be blank")
	}
	if a.OwnerId == 0 {
		return fmt.Errorf("invalid owner id: %d", a.OwnerId)
	}
	return nil
}
