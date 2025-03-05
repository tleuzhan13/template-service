package repository

import "template-service/internal/model"

type (
	userDB struct {
		ID         uint64 `bson:"_id"`
		FirstName  string `bson:"first_name"`
		SecondName string `bson:"second_name"`
	}
)

func (u *userDB) Set(user *model.User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.SecondName = user.SecondName
}

func (u *userDB) ToModel() *model.User {
	return &model.User{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
	}
}
