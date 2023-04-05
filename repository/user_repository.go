package repository

import (
	"project-tiga/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Add(newUser model.User) error {
	tx := ur.db.Create(&newUser)
	return tx.Error
}

func (ur *UserRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	tx := ur.db.First(&user, "email = ?", email)
	return user, tx.Error
}

func (ur *UserRepository) GetByID(userID string) (model.User, error) {
	var user model.User
	tx := ur.db.First(&user, "id = ?", userID)
	return user, tx.Error
}
