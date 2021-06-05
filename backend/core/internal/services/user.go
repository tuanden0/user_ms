package services

import (
	"context"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/models"
	"user_ms/backend/core/internal/repository"
)

type UserService interface {
	api.UserAPIServer
}

type userService struct {
	repo repository.UserRepository
	api.UnimplementedUserAPIServer
}

func Init(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Ping(ctx context.Context, in *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{
		Message: "Pong",
	}, nil
}

func (s *userService) Create(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {

	u := &models.User{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
	}

	hash, err := u.HashPassword()

	if err != nil {
		return nil, err
	}

	u.Password = hash

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return &api.CreateUserResponse{
		Message: "create user success",
	}, nil
}
