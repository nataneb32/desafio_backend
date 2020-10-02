package gorm_checkin

import (
	"fmt"
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/checkin"
)

type CheckInRepo struct {
	DB *gorm.DB
}

func CreateCheckInRepo(db *gorm.DB) *CheckInRepo {
	return &CheckInRepo{
		DB: db,
	}
}

func (cr *CheckInRepo) CreateCheckIn(ci *checkin.CheckIn) error {
	return cr.DB.Model(&checkin.CheckIn{}).Create(ci).Error
}
func (cr *CheckInRepo) GetCheckIn(ci *checkin.CheckIn) error {
	return cr.DB.Model(&checkin.CheckIn{}).Where(ci).First(ci).Error
}
func (cr *CheckInRepo) UpdateCheckIn(id uint, ci *checkin.CheckIn) error {
	return cr.DB.Model(&checkin.CheckIn{}).Where("id = ?", id).Updates(ci).Error
}

func (cr *CheckInRepo) GetAllCheckIn(ci *checkin.CheckIn) (error, []checkin.CheckIn) {
	var result []checkin.CheckIn
	err := cr.DB.Model(&checkin.CheckIn{}).Where(ci).Find(&result).Error
	return err, result
}
func (cr *CheckInRepo) GetNewestCheckInOf(guestId uint) (error, checkin.CheckIn) {
	var result checkin.CheckIn
	err := cr.DB.Model(&checkin.CheckIn{}).
		Order("data_entrada desc").
		Where("hospede = ?", guestId).
		Where("data_saida is not NULL").
		Limit(1).
		Find(&result).Error
	fmt.Println(result)
	return err, result
}
