package gorm_checkin

import (
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/checkins"
)

type CheckInSchema struct {
	gorm.Model
}

type CheckInRepo struct {
	DB *gorm.DB
}

func CreateCheckInRepo(db *gorm.DB) checkins.CheckInRepo {
	return &CheckInRepo{
		DB: db,
	}
}

func (cr *CheckInRepo) CreateCheckIn(ci *checkins.CheckIn) error {
	return cr.DB.Model(&CheckInSchema{}).Create(ci).Error
}
