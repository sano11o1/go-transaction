package main

import (
	"github.com/sano11o1/go-transaction/model"
	"github.com/sano11o1/go-transaction/repository"
	"github.com/sano11o1/go-transaction/usecase"
	"gorm.io/gorm"
)

func main() {
	db := initDB()
	userRepostiry := repository.NewUserRepositoryImpl(db)
	registerUserUsecase := usecase.NewRegisterUserUsecase(userRepostiry)
	newUser := model.User{
		Name: "sano11o1",
	}
	if err := registerUserUsecase.Execute(newUser); err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	var db *gorm.DB
	return db
}
