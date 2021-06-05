package repository

import "user_ms/backend/core/internal/models"

type UserRepository interface {
	Create(u *models.User) error
	Retrieve(id string) (*models.User, error)
	Update(id string, in models.User) (*models.User, error)
	Delete(id string) error
	List() ([]*models.User, error)
}
