package usecase

import (
	"github.com/sano11o1/go-transaction/model"
	"github.com/sano11o1/go-transaction/repository"
)

type RegisterUserUsecase struct {
	userRepository repository.IUserReopsitory
}

func NewRegisterUserUsecase(userRepository repository.IUserReopsitory) *RegisterUserUsecase {
	return &RegisterUserUsecase{
		userRepository: userRepository,
	}
}

func (u *RegisterUserUsecase) Execute(user model.User) error {
	// TODO 実装
	return nil
}
