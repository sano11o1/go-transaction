package repository

import (
	"gorm.io/gorm"
)

type IBaseRepository interface {
	Atmoic(fn func(IBaseRepository) error) error
	GetUserRepository() IUserRepository
	GetActivityLogRepository() IActivityLogRepository
}

type BaseRepositoryImpl struct {
	db *gorm.DB
}

func NewBaseRepositoryImpl(db *gorm.DB) IBaseRepository {
	return &BaseRepositoryImpl{
		db: db,
	}
}

func (r *BaseRepositoryImpl) Atmoic(fn func(IBaseRepository) error) error {
	// r.dbするにはmain.goでdbを初期化しBaseRepositoryImplを生成する際に渡す必要がある
	return r.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewBaseRepositoryImpl(tx))
	})
}

func (r *BaseRepositoryImpl) GetUserRepository() IUserRepository {
	// NewUserRepositoryImplの引数にはBaseRepositoryImplのDBを渡す
	return NewUserRepositoryImpl(r.db)
}

func (r *BaseRepositoryImpl) GetActivityLogRepository() IActivityLogRepository {
	// NewUserRepositoryImplの引数にはBaseRepositoryImplのDBを渡す
	return NewActivityLogRepositoryImpl(r.db)
}
