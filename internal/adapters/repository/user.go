package repository

import (
	"context"
	"errors"
	"template-service/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	conn       *mongo.Database
	collection string
}

const usersCollection = "users"

func NewUser(conn *mongo.Database) *UserRepo {
	return &UserRepo{
		conn:       conn,
		collection: usersCollection,
	}
}

func (s *UserRepo) Upsert(ctx context.Context, user *model.User) error {
	var err error
	var userToDB userDB
	userToDB.Set(user)

	opt := options.Update().SetUpsert(true)
	_, err = s.conn.Collection(s.collection).
		UpdateOne(ctx, primitive.M{"_id": userToDB.ID}, primitive.M{"$set": userToDB}, opt)

	if err != nil {

		return err
	}

	return nil
}

func (s *UserRepo) Get(ctx context.Context, userID uint64) (*model.User, error) {
	var err error
	var result userDB

	err = s.conn.Collection(s.collection).
		FindOne(ctx, primitive.M{"_id": userID, "deleted": false}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, model.ErrNotFound
		}

		return nil, err
	}

	return result.ToModel(), nil
}

func (s *UserRepo) GetAll(ctx context.Context) ([]*model.User, error) {

	return nil, nil
}

func (s *UserRepo) Delete(ctx context.Context, userID uint64) error {

	return nil
}
