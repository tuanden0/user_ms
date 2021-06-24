package repository

import (
	"user_ms/backend/core/internal/models"
)

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

func (p Pagination) GetLimit() uint32 {

	if p.Limit <= 0 {
		return 5
	}

	if p.Limit > 100 {
		return 100
	}

	return p.Limit
}

func (p Pagination) GetPage() uint32 {

	if p.Page <= 0 {
		return 1
	}

	return p.Page
}

type Sort struct {
	Key   string
	IsASC bool
}

func (s Sort) GetKey() string {
	switch s.Key {
	case "id", "username", "email", "role":
		return s.Key
	default:
		return "id"
	}
}

func (s Sort) GetIsASC() string {
	if s.IsASC {
		return "ASC"
	}
	return "DESC"
}

type Filter struct {
	Key    string
	Value  string
	Method string
}

func (f Filter) GetKey() string {
	switch f.Key {
	case "id", "username", "email", "role":
		return f.Key
	default:
		return "id"
	}
}

func (f Filter) GetValue() string {
	return f.Value
}

func (f Filter) GetMethod() string {
	switch f.Method {
	case ">", ">=", "<", "<=", "=":
		return f.Method
	default:
		return "="
	}
}
