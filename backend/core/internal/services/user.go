package services

import (
	"context"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/repository"
	"user_ms/backend/core/internal/util"
	"user_ms/backend/core/internal/validators"
)

type UserService interface {
	api.UserAPIServer
}

type userService struct {
	repo repository.UserRepository
	api.UnimplementedUserAPIServer
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Ping(ctx context.Context, in *api.PingRequest) (*api.PingResponse, error) {

	// Validate User access role
	_, err := repository.ValidateAcessRole(ctx, "all")
	if err != nil {
		return nil, err
	}

	return &api.PingResponse{
		Message: "Pong",
	}, nil
}

func (s *userService) Create(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {

	// Validate CreateUserRequest
	if err := validators.ValidateCreateUserRequest(ctx, in); err != nil {
		return nil, err
	}

	// Mapping input to model
	u, err := util.MapCreateUserRequest(ctx, in)
	if err != nil {
		return nil, err
	}

	// Create user
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}

	// Response to client
	return util.MapCreateUserResponse("create user success"), nil
}

func (s *userService) Retrieve(ctx context.Context, in *api.RetrieveUserRequest) (*api.User, error) {

	// Validate RetrieveUserRequest
	if err := validators.ValidateRetrieveUserRequest(ctx, in); err != nil {
		return nil, err
	}

	// Mapping input to user id
	id := util.MapUserId(in.GetId())

	// Get user from user id
	u, err := s.repo.Retrieve(id)
	if err != nil {
		return nil, err
	}

	// Response to client
	return util.MapRetrieveUserResponse(u), nil
}

func (s *userService) Update(ctx context.Context, in *api.UpdateUserRequest) (*api.User, error) {

	// Validate UpdateUserRequest
	if err := validators.ValidateUpdateUserRequest(ctx, in); err != nil {
		return nil, err
	}

	// Mapping user input to model
	uInput, err := util.MapUpdateUserRequest(ctx, in)
	if err != nil {
		return nil, err
	}

	// Update user
	u, err := s.repo.Update(uInput.GetStringID(), *uInput)
	if err != nil {
		return nil, err
	}

	// Return to client
	return util.MapUpdateUserResponse(u), nil
}

func (s *userService) Delete(ctx context.Context, in *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {

	// Validate DeleteUserRequest
	if err := validators.ValidateDeleteUserRequest(ctx, in); err != nil {
		return nil, err
	}

	// Mapping input
	id := util.MapUserId(in.GetId())
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}

	// Respose to client
	return util.MapDeleteUserResponse("delete user success"), nil
}

func (s *userService) List(ctx context.Context, in *api.ListUserRequest) (*api.ListUserResponse, error) {

	// Validate ListUserRequest
	if err := validators.ValidateListUserRequest(ctx, in); err != nil {
		return nil, err
	}

	// Mapping input to models
	pagination, sort, filters := util.MapListUserRequest(in)

	// Get list users
	users, err := s.repo.List(pagination, sort, filters)
	if err != nil {
		return nil, err
	}

	// Response to client
	return util.MapListUserResponse(users), nil
}
