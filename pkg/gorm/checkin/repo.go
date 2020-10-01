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

func CreateCheckInRepo(db *gorm.DB) *CheckInRepo {
	return &CheckInRepo{
		DB: db,
	}
}

func (cr *CheckInRepo) CreateCheckIn(ci *checkins.CheckIn) error {
	return cr.DB.Model(&CheckInSchema{}).Create(ci).Error
}
func (cr *CheckInRepo) GetCheckIn(ci *checkins.CheckIn) error {
	return cr.DB.Model(&CheckInSchema{}).First(ci).Error
}
