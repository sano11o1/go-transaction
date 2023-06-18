package repository

import (
	"github.com/sano11o1/go-transaction/model"
	"gorm.io/gorm"
)

type IUserReopsitory interface {
	AddUser(model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) IUserReopsitory {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) AddUser(user model.User) error {
	return r.db.Create(&user).Error
}
