//go:build !unit
// +build !unit

package entity

import (
	"errors"
	"strings"
)

var (
	ErrMissingUsername = errors.New("username cannot be empty")
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (u *User) Validate() error {
	if strings.TrimSpace(u.Name) == "" {
		return ErrMissingUsername
	}
	return nil
}
