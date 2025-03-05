package usecase

import (
	"context"
	"template-service/internal/model"
	"template-service/internal/ports"
)

type UserUseCase struct {
	repo ports.UserRepo
}

func NewUserUseCase(r ports.UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) Upsert(ctx context.Context, user *model.User) error {
	if user == nil {
		return ErrEmptyUser
	}

	return uc.repo.Upsert(ctx, user)
}

func (uc *UserUseCase) Get(ctx context.Context, userID uint64) (*model.User, error) {
	if userID == 0 {
		return &model.User{}, ErrEmptyUserID
	}

	user, err := uc.repo.Get(ctx, userID)
	if err != nil {
		return &model.User{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetAll(ctx context.Context) ([]*model.User, error) {
	users, err := uc.repo.GetAll(ctx)
	if err != nil {
		return []*model.User{}, err
	}

	return users, nil
}

func (uc *UserUseCase) Delete(ctx context.Context, userID uint64) error {
	if userID == 0 {
		return ErrEmptyUserID
	}

	return uc.repo.Delete(ctx, userID)
}
