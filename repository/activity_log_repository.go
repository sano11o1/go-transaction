package repository

import (
	"github.com/sano11o1/go-transaction/entity"
	"gorm.io/gorm"
)

type IActivityLogRepository interface {
	AddActivityLog(log entity.ActivityLog) error
}

type ActivityLogRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityLogRepositoryImpl(db *gorm.DB) IActivityLogRepository {
	return &ActivityLogRepositoryImpl{
		db: db,
	}
}
func (r *ActivityLogRepositoryImpl) AddActivityLog(log entity.ActivityLog) error {
	return r.db.Create(&log).Error
	//return errors.New("エラー")
}
