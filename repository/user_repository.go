package repository

import (
	"errors"
	"fmt"

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

type UserRepositoryTestImpl struct {
	//addUserFunc func(user entity.User) error
}

func NewUserRepositoryTestImpl() IUserRepository {
	return &UserRepositoryTestImpl{
		//addUserFunc: addUserFunc,
	}
}

func (r *UserRepositoryTestImpl) AddUser(user entity.User) error {
	fmt.Println("========= AddUser Test Method =========")
	return errors.New("failed to add user") //r.addUserFunc(user)
}
