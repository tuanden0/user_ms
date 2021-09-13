package validators

import (
	"context"
	"fmt"
	"user_ms/backend/core/api"
	"user_ms/backend/core/internal/repository"
)

func ValidateUserLoginRequest(ctx context.Context, in *api.UserLoginRequest) error {

	// Validate User access role and get current user
	userClaim, accessErr := repository.ValidateAcessRole(ctx, "all")
	if accessErr != nil {
		return accessErr
	}

	// Only non-login user can access
	if !userClaim.IsAnonymous() {
		return fmt.Errorf("already logged in")
	}

	return nil
}
