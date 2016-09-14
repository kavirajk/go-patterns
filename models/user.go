package models

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

type User struct {
	Model
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
	Friends   []User `json:"-"`
}

func (u User) String() string {
	return fmt.Sprintf("<%d: %s>", u.Id, u.Username)
}

func (u *User) Valid() error {
	if strings.TrimSpace(u.Username) == "" || len(u.Username) < 4 {
		return fmt.Errorf("username too short")
	}
	if strings.TrimSpace(u.Password) == "" || len(u.Password) < 8 {
		return fmt.Errorf("password too short")
	}
	return nil
}

func (u *User) BeforeSave() error {
	u.Username = strings.TrimSpace(strings.ToLower(u.Username))

	// Hash the password
	hasher := sha1.New()
	if _, err := hasher.Write([]byte(u.Password)); err != nil {
		return err
	}
	u.Password = hex.EncodeToString(hasher.Sum(nil))
	return nil
}

type UserFriends struct {
	UserId   uint
	FriendId uint
}
