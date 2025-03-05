// nolint
package model

import (
	"testing"
)

func TestUserModel(t *testing.T) *User {
	return &User{
		FirstName:  "TestFirstName",
		SecondName: "TestSecondName",
		ID:         1,
	}
}
