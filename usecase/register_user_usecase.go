package usecase

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sano11o1/go-transaction/entity"
	"github.com/sano11o1/go-transaction/repository"
)

type RegisterUserUsecase struct {
	baseRepository repository.IBaseRepository
}

func NewRegisterUserUsecase(baseRepository repository.IBaseRepository) *RegisterUserUsecase {
	return &RegisterUserUsecase{
		baseRepository: baseRepository,
	}
}

func (u *RegisterUserUsecase) Execute(user entity.User) error {

	atomicBlock := func(r repository.IBaseRepository) error {
		// ユーザーを作成
		userRepo := r.GetUserRepository()
		if err := userRepo.AddUser(user); err != nil {
			return err
		}
		// ログを作成
		logRepo := r.GetActivityLogRepository()
		log := entity.ActivityLog{
			ID:           uuid.New(),
			ActivityType: "user_register",
			CreatedAt:    time.Now(),
		}
		if err := logRepo.AddActivityLog(log); err != nil {
			return err
		}

		// 外部サービスにUserを作成
		req, err := http.NewRequest("POST", "https://example.com/user", nil)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		client := new(http.Client)
		_, err = client.Do(req)
		if err != nil {
			fmt.Println("failed to create user")
			// 外部サービスとの通信が失敗した場合ロールバックする
			return err
		}
		fmt.Println("success to create user")
		return nil
	}
	err := u.baseRepository.Atmoic(atomicBlock)
	return err
}
