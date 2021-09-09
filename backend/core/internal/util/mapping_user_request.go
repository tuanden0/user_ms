package util

import (
	"context"
	"fmt"
	"strconv"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/models"
	"user_ms/backend/core/internal/repository"
)

func MapCreateUserRequest(ctx context.Context, in *api.CreateUserRequest) (*models.User, error) {

	// Mapping user input to user model
	u := &models.User{
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
		Role:     in.GetRole(),
	}

	// Handle hash pasword
	hash, err := u.HashPassword()
	if err != nil {
		return nil, err
	}
	u.Password = hash

	return u, nil
}

func MapCreateUserResponse(message string) *api.CreateUserResponse {
	return &api.CreateUserResponse{
		Message: message,
	}
}

func MapRetrieveUserResponse(u *models.User) *api.User {
	return &api.User{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
	}
}

func MapUpdateUserRequest(ctx context.Context, in *api.UpdateUserRequest) (*models.User, error) {

	// Get user role
	userClaim := repository.ParseUsersOrNilFromCTX(ctx)
	if userClaim == nil {
		return nil, fmt.Errorf("unable to get user info")
	}

	u := &models.User{
		ID:       in.GetId(),
		Username: in.GetUsername(),
		Password: in.GetPassword(),
		Email:    in.GetEmail(),
	}

	if userClaim.GetRole() == "admin" {
		u.Role = in.GetRole()
	}

	if u.GetPassWord() != "" {
		hash, err := u.HashPassword()
		if err != nil {
			return nil, err
		}

		u.Password = hash
	}

	return u, nil
}

func MapUpdateUserResponse(u *models.User) *api.User {
	return &api.User{
		Id:       u.GetID(),
		Username: u.GetUserName(),
		Email:    u.GetEmail(),
	}
}

func MapUserId(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func MapDeleteUserResponse(message string) *api.DeleteUserResponse {
	return &api.DeleteUserResponse{
		Message: message,
	}
}

func MapListUserRequest(in *api.ListUserRequest) (*repository.Pagination, *repository.Sort, []*repository.Filter) {

	inputSort := in.GetSort()
	inputFilters := in.GetFilters()
	inputPagination := in.GetPagination()

	sort := repository.NewSort(inputSort.GetKey(), inputSort.GetIsAsc())

	pagination := repository.NewPagination(inputPagination.GetLimit(), inputPagination.GetPage())

	filters := make([]*repository.Filter, 0)
	for _, f := range inputFilters {
		filters = append(filters, repository.NewFilter(f.GetKey(), f.GetValue(), f.GetMethod()))
	}

	return pagination, sort, filters
}

func MapListUserResponse(users []*models.User) *api.ListUserResponse {
	res := make([]*api.User, 0)
	for _, u := range users {
		res = append(res, &api.User{
			Id:       u.GetID(),
			Username: u.GetUserName(),
			Email:    u.GetEmail(),
		})
	}

	return &api.ListUserResponse{
		Users: res,
	}
}
