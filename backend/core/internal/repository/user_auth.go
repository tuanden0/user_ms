package repository

import "user_ms/backend/core/internal/models"

type UserAuthRepository interface {
	Login(username string) (*models.User, error)
	GenerateJWTToken(*models.User) (string, error)
}
