package services

import (
	"context"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserAuthenService interface {
	api.AuthenServiceServer
}

type userAuthenService struct {
	repo repository.UserAuthRepository
	api.UnimplementedAuthenServiceServer
}

func NewUserAuthenService(repo repository.UserAuthRepository) UserAuthenService {
	return &userAuthenService{repo: repo}
}

func (s *userAuthenService) Ping(ctx context.Context, in *api.PingRequest) (*api.PingResponse, error) {
	return &api.PingResponse{
		Message: "Pong",
	}, nil
}

func (s *userAuthenService) Login(ctx context.Context, in *api.UserLoginRequest) (*api.UserLoginResponse, error) {

	username := in.GetUsername()
	password := in.GetPassword()

	u, err := s.repo.Login(username)
	if err != nil {
		return nil, err
	}

	if !u.CheckPassword(password) || u == nil {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := s.repo.GenerateJWTToken(u)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	return &api.UserLoginResponse{
		AccessToken: token,
	}, nil
}
