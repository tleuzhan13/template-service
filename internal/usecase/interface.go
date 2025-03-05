//go:generate mockgen -source=interface.go -destination=../../mocks/mockusecase/mockusecase.go -package=mockusecase
package usecase

import (
	"context"

	"template-service/internal/model"
)

type User interface {
	Upsert(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id uint64) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
	Delete(ctx context.Context, id uint64) error
}
