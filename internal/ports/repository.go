//go:generate mockgen -source=repository.go -destination=../../mocks/mockrepository/mockrepository.go -package=mockrepository
package ports

import (
	"context"

	"template-service/internal/model"
)

type UserRepo interface {
	Upsert(ctx context.Context, user *model.User) error
	Get(ctx context.Context, userID uint64) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	Delete(ctx context.Context, userID uint64) error
}
