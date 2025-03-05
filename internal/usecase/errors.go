package usecase

import (
	"errors"
)

var (
	ErrEmptyUser   = errors.New("user data is empty")
	ErrEmptyUserID = errors.New("user id is invalid")
)
