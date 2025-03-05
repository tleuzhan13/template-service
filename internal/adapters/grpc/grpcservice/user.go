package grpcservice

import (
	"context"

	// TODO replace with import path to grpc generated code
	base "/gen/grpc-repository/template-service/base"
	evt "/gen/grpc-repository/template-service/events"
	svc "/gen/grpc-repository/template-service/service"

	"template-service/internal/usecase"
)

type UserService struct {
	svc.HubServiceServer
	uc usecase.User
}

func NewUserService(uc usecase.User) *UserService {
	return &UserService{
		uc: uc,
	}
}

func (s *UserService) GetUser(
	ctx context.Context,
	req *svc.GetUserRequest) (*svc.GetUserResponse, error) {

	user, err := s.uc.Get(ctx, req.GetID())
	if err != nil {
		return &svc.GetUserResponse{
			Error: &base.Error{
				Message: err.Error(),
			}}, err
	}

	return &svc.GetUserResponse{
		ID:              user.ID,
		user.FirstName:  user.FirstName,
		user.SecondName: user.SecondName,
	}, nil
}

func (s *UserService) UpsertUserGetUser(
	ctx context.Context,
	req *svc.UpsertUserGetUserRequest) (*svc.UpsertUserGetUserResponse, error) {

	return nil, nil
}

func (s *UserService) ListUserGetUsers(
	ctx context.Context,
	req *svc.ListUserGetUsersRequest) (*svc.ListUserGetUsersResponse, error) {

	return nil, nil
}

func (s *UserService) DeleteUserGetUser(
	ctx context.Context,
	req *svc.DeleteUserGetUserRequest,
) (
	*svc.DeleteUserGetUserResponse,
	error,
) {
	var err error

	return nil, err
}
