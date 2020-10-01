package gorm_guest

import (
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/gorm/checkin"
	"nataneb32.live/hospedagem/pkg/guests"
)

type GuestSchema struct {
	ID        uint `gorm:"primaryKey"`
	Nome      string
	Documento string
	Telefone  string
	CheckIns  []gorm_checkin.CheckInSchema `gorm:"foreignKey:Hospede"`
}

type GuestRepo struct {
	DB *gorm.DB
}

// Return The First Guest that matches with g
func (gs *GuestRepo) GetGuest(g *guests.Guest) error {
	err := gs.DB.Model(&GuestSchema{}).Where(&g).First(&g).Error
	return err
}

// Creates a new Guest
func (gs *GuestRepo) CreateGuest(g *guests.Guest) error {
	// Setting to nil, because cannot create a guest with checkins.
	g.CheckIns = nil
	err := gs.DB.Model(&GuestSchema{}).Create(g).Error
	return err
}

// Create and return a GuestRepo
func CreateGuestRepo(db *gorm.DB) guests.GuestRepo {
	return &GuestRepo{
		DB: db,
	}
}
