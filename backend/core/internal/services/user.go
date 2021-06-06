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

func (s *userService) Retrieve(ctx context.Context, in *api.RetrieveUserRequest) (*api.User, error) {

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

	uInput := models.User{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
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
	id := strconv.FormatUint(uint64(in.GetId()), 10)
	if err := s.repo.Delete(id); err != nil {
		return nil, err
	}

	return &api.DeleteUserResponse{
		Message: "delete user success",
	}, nil
}

func (s *userService) List(ctx context.Context, in *api.ListUserRequest) (*api.ListUserResponse, error) {

	inputSort := in.GetSort()
	inputFilters := in.GetFilters()
	inputPagination := in.GetPagination()

	sort := repository.NewSort("id", "ASC")
	if inputSort != nil {
		sort.Key = inputSort.Key
		if !inputSort.IsAsc {
			sort.IsASC = "DESC"
		}

	}

	pagination := repository.NewPagination(5, 1)
	if inputPagination != nil {
		pagination.Limit = inputPagination.Limit
		pagination.Page = inputPagination.Page
	}

	filters := make([]*repository.Filter, 0)
	if filters != nil {
		for _, f := range inputFilters {
			filters = append(filters, repository.NewFilter(f.Key, f.Method, f.Value))
		}
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
