package main

import (
	"fmt"

	"github.com/sano11o1/go-transaction/entity"
	"github.com/sano11o1/go-transaction/repository"
	"github.com/sano11o1/go-transaction/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	baseRepostiry := repository.NewBaseRepositoryImpl(db)
	registerUserUsecase := usecase.NewRegisterUserUsecase(baseRepostiry)
	newUser := entity.User{
		Name: "sano11o1",
	}
	if err := registerUserUsecase.Execute(newUser); err != nil {
		fmt.Println("failed to register user", err.Error())
		panic(err)
	}
}

func initDB() (*gorm.DB, error) {
	host := "127.0.0.1"
	port := "5431"
	user := "postgres"
	dbName := "postgres"
	password := "passw0rd"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbName, port)
	var err error
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
		return nil, err
	}

	return db, nil
}
