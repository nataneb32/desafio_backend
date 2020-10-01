package gorm_checkin

import (
	"fmt"
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/checkins"
	"time"
)

type CheckInSchema struct {
	ID               uint      `gorm:"primaryKey"`
	Hospede          uint      ``
	DataEntrada      time.Time ``
	DataSaida        time.Time ``
	AdicionalVeiculo bool      ``
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
	fmt.Println(ci)
	return cr.DB.Model(&CheckInSchema{}).Where(ci).First(ci).Error
}
