package repository

import (
	"user_ms/backend/core/internal/models"

	"gorm.io/gorm"
)

type UserMng struct {
	db *gorm.DB
}

func NewUserMng(db *gorm.DB) UserRepository {
	return &UserMng{db: db}
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

func (m *UserMng) List() ([]*models.User, error) {

	users := make([]*models.User, 0)

	if err := m.db.
		Find(&users).
		Error; err != nil {
		return users, err
	}

	return users, nil
}
