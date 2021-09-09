package validators

import (
	"context"
	"fmt"
	"log"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/repository"
)

func ValidateCreateUserRequest(ctx context.Context, in *api.CreateUserRequest) error {

	// Validate User access role
	_, accessErr := repository.ValidateAcessRole(ctx, "admin")
	if accessErr != nil {
		return accessErr
	}
	return nil
}

func ValidateRetrieveUserRequest(ctx context.Context, in *api.RetrieveUserRequest) error {

	// Validate User access role
	userClaim, accessErr := repository.ValidateAcessRole(ctx, "admin", "user")
	if accessErr != nil {
		return accessErr
	}

	// Validate user id
	if userClaim.GetID() != in.GetId() {
		if userClaim.GetRole() != "admin" {
			log.Printf("%v try to get user id %v", userClaim.GetUserName(), in.GetId())
			return fmt.Errorf("access denied")
		}
	}

	return nil
}

func ValidateUpdateUserRequest(ctx context.Context, in *api.UpdateUserRequest) error {

	// Validate User access role and get current user
	userClaim, accessErr := repository.ValidateAcessRole(ctx, "admin", "user")
	if accessErr != nil {
		return accessErr
	}

	// Validate user id
	if userClaim.GetID() != in.GetId() {
		if userClaim.GetRole() != "admin" {
			log.Printf("%v try to update user id %v", userClaim.GetUserName(), in.GetId())
			return fmt.Errorf("access denied")
		}
	}

	return nil
}

func ValidateDeleteUserRequest(ctx context.Context, in *api.DeleteUserRequest) error {

	// Validate User access role and get current user
	userClaim, accessErr := repository.ValidateAcessRole(ctx, "admin", "user")
	if accessErr != nil {
		return accessErr
	}

	// Validate user id
	if userClaim.GetID() != in.GetId() {
		if userClaim.GetRole() != "admin" {
			log.Printf("%v try to delete user id %v", userClaim.GetUserName(), in.GetId())
			return fmt.Errorf("access denied")
		}
	}

	return nil
}

func ValidateListUserRequest(ctx context.Context, in *api.ListUserRequest) error {

	// Validate User access role
	_, accessErr := repository.ValidateAcessRole(ctx, "admin")
	if accessErr != nil {
		return accessErr
	}

	return nil
}
