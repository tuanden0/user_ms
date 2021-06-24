package services

import (
	"context"
	"strconv"
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

	// Validate User access role
	_, err := repository.ValidateAcessRole(ctx, "admin")
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
		Role:     in.GetRole(),
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

func (s *userService) Retrieve(ctx context.Context, in *api.RetrieveUserRequest) (*api.User, error) {

	// Validate User access role
	_, err := repository.ValidateAcessRole(ctx, "admin", "user")
	if err != nil {
		return nil, err
	}

	id := strconv.FormatUint(uint64(in.GetId()), 10)
	u, err := s.repo.Retrieve(id)
	if err != nil {
		return nil, err
	}

	return &api.User{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
	}, nil
}

func (s *userService) Update(ctx context.Context, in *api.UpdateUserRequest) (*api.User, error) {

	// Validate User access role
	userClaim, err := repository.ValidateAcessRole(ctx, "admin", "user")
	if err != nil {
		return nil, err
	}

	uInput := models.User{}

	if userClaim.Role == "admin" {
		uInput = models.User{
			Username: in.GetUsername(),
			Password: in.GetPassword(),
			Email:    in.GetEmail(),
			Role:     in.GetRole(),
		}
	} else {
		uInput = models.User{
			Username: in.GetUsername(),
			Password: in.GetPassword(),
			Email:    in.GetEmail(),
		}
	}

	hash, err := uInput.HashPassword()

	if err != nil {
		return nil, err
	}

	uInput.Password = hash

	id := strconv.FormatUint(uint64(in.GetId()), 10)

	u, err := s.repo.Update(id, uInput)
	if err != nil {
		return nil, err
	}

	return &api.User{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
	}, nil
}

func (s *userService) Delete(ctx context.Context, in *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {

	// Validate User access role
	_, err := repository.ValidateAcessRole(ctx, "admin", "user")
	if err != nil {
		return nil, err
	}

	id := strconv.FormatUint(uint64(in.GetId()), 10)
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}

	return &api.DeleteUserResponse{
		Message: "delete user success",
	}, nil
}

func (s *userService) List(ctx context.Context, in *api.ListUserRequest) (*api.ListUserResponse, error) {

	// Validate User access role
	_, err := repository.ValidateAcessRole(ctx, "admin")
	if err != nil {
		return nil, err
	}

	inputSort := in.GetSort()
	inputFilters := in.GetFilters()
	inputPagination := in.GetPagination()

	sort := repository.NewSort(inputSort.GetKey(), inputSort.GetIsAsc())

	pagination := repository.NewPagination(inputPagination.GetLimit(), inputPagination.GetPage())

	filters := make([]*repository.Filter, 0)
	for _, f := range inputFilters {
		filters = append(filters, repository.NewFilter(f.GetKey(), f.GetValue(), f.GetMethod()))
	}

	uList, err := s.repo.List(pagination, sort, filters)
	if err != nil {
		return nil, err
	}

	res := make([]*api.User, 0)
	for _, u := range uList {
		res = append(res, &api.User{
			Id:       u.GetID(),
			Username: u.GetUserName(),
			Email:    u.GetEmail(),
		})
	}

	return &api.ListUserResponse{
		Users: res,
	}, nil
}
