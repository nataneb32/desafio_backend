package gorm_checkin

import (
	"fmt"
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
func (cr *CheckInRepo) UpdateCheckIn(id uint, ci *checkins.CheckIn) error {
	return cr.DB.Model(&checkins.CheckIn{}).Where("id = ?", id).Updates(ci).Error
}

func (cr *CheckInRepo) GetAllCheckIn(ci *checkins.CheckIn) (error, []checkins.CheckIn) {
	var result []checkins.CheckIn
	err := cr.DB.Model(&checkins.CheckIn{}).Where(ci).Find(&result).Error
	return err, result
}
func (cr *CheckInRepo) GetNewestCheckInOf(guestId uint) (error, checkins.CheckIn) {
	var result checkins.CheckIn
	err := cr.DB.Model(&checkins.CheckIn{}).
		Order("data_entrada desc").
		Where("hospede = ?", guestId).
		Where("data_saida is not NULL").
		Limit(1).
		Find(&result).Error
	fmt.Println(result)
	return err, result
}
