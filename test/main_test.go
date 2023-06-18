package test

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/sano11o1/go-transaction/entity"
	"github.com/sano11o1/go-transaction/repository"
	"github.com/sano11o1/go-transaction/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(t *testing.T) {
	db, err := initDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepositoryTestImpl()

	baseRepo := repository.NewTestBaseRepositoryImpl(db, userRepo)

	usecase := usecase.NewRegisterUserUsecase(baseRepo)
	err = usecase.Execute(entity.User{Name: "aiueo"})
	if err == nil {
		t.Errorf("エラーを返さなかった")
		return
	}
	assert.Equal(t, err.Error(), "failed to add user")
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
