package repository

import (
	"github.com/sano11o1/go-transaction/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	AddUser(entity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) IUserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) AddUser(user entity.User) error {
	return r.db.Create(&user).Error
}
