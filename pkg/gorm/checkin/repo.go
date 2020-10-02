package gorm_checkin

import (
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/checkins"
)

type CheckInRepo struct {
	DB *gorm.DB
}

func CreateCheckInRepo(db *gorm.DB) *CheckInRepo {
	return &CheckInRepo{
		DB: db,
	}
}

func (cr *CheckInRepo) CreateCheckIn(ci *checkins.CheckIn) error {
	return cr.DB.Model(&checkins.CheckIn{}).Create(ci).Error
}
func (cr *CheckInRepo) GetCheckIn(ci *checkins.CheckIn) error {
	return cr.DB.Model(&checkins.CheckIn{}).Where(ci).First(ci).Error
}
