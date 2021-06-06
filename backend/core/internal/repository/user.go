package repository

import "user_ms/backend/core/internal/models"

type UserRepository interface {
	Create(u *models.User) error
	Retrieve(id string) (*models.User, error)
	Update(id string, in models.User) (*models.User, error)
	Delete(id string) error
	List(pagination *Pagination, sort *Sort, filters []*Filter) ([]*models.User, error)
}

type Pagination struct {
	Limit uint32
	Page  uint32
}

type Sort struct {
	Key   string
	IsASC string
}

type Filter struct {
	Key    string
	Value  string
	Method string
}
