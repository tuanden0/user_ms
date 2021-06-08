package repository

import (
	"fmt"
	"user_ms/backend/core/internal/models"

	"gorm.io/gorm"
)

type UserMng struct {
	db *gorm.DB
}

func NewUserMng(db *gorm.DB) UserRepository {
	return &UserMng{db: db}
}

func NewSort(key string, is_asc string) *Sort {
	return &Sort{
		Key:   key,
		IsASC: is_asc,
	}
}

func NewFilter(key string, value string, method string) *Filter {
	return &Filter{
		Key:    key,
		Value:  value,
		Method: method,
	}
}

func NewPagination(limit uint32, page uint32) *Pagination {
	return &Pagination{
		Limit: limit,
		Page:  page,
	}
}

// Implement UserRepository functions

func (m *UserMng) Create(u *models.User) error {
	if err := m.db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (m *UserMng) Retrieve(id string) (*models.User, error) {

	user := models.User{}

	if err := m.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserMng) Update(id string, in models.User) (*models.User, error) {

	user := models.User{}

	if err := m.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if err := m.db.Model(&user).Omit("id").Updates(&in).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserMng) Delete(id string) error {

	if err := m.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (m *UserMng) List(pagination *Pagination, sort *Sort, filters []*Filter) ([]*models.User, error) {

	users := make([]*models.User, 0)

	qs := m.db.
		Order(fmt.Sprintf("%v %v", sort.Key, sort.IsASC)).
		Limit(int(pagination.Limit)).
		Offset(int(pagination.Limit) * (int(pagination.Page) - 1))

	for _, f := range filters {
		qs = qs.Where(fmt.Sprintf("%v %v ?", f.Key, f.Method), f.Value)
	}

	if err := qs.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
