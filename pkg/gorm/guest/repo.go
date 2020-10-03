package gorm_guest

import (
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/checkin"
	"nataneb32.live/hospedagem/pkg/guests"
)

type GuestRepo struct {
	DB *gorm.DB
}

// Return The First Guest that matches with g
func (gs *GuestRepo) GetGuest(g *guests.Guest) error {
	err := gs.DB.Model(&guests.Guest{}).Where(&g).Preload("CheckIns").First(&g).Error
	return err
}

// Creates a new Guest
func (gs *GuestRepo) CreateGuest(g *guests.Guest) error {
	// Setting id to 0, because the id is automatically generated.
	g.ID = 0
	// Setting to nil, because cannot create a guest with checkins.
	g.CheckIns = nil
	err := gs.DB.Model(&guests.Guest{}).Create(g).Error
	return err
}

func (gs *GuestRepo) UpdateGuest(id uint, g *guests.Guest) error {
	return gs.DB.Model(&guests.Guest{}).Where("id = ?", id).Updates(g).Error
}

// Search Guest
func (gs *GuestRepo) SearchGuest(query guests.GuestQuery) struct {
	Guests     []guests.Guest
	TotalPages uint
} {
	// Default Value to Page
	if query.Page == 0 {
		query.Page = 1
	}
	// Default Value to Limit
	if query.Limit == 0 {
		query.Limit = 10
	}

	var result struct {
		Guests     []guests.Guest
		TotalPages uint
	}

	var count int64
	gs.DB.Model(&guests.Guest{}).
		Where("telefone LIKE ?", "%"+query.Telefone+"%").
		Where("nome LIKE ?", "%"+query.Nome+"%").
		Where("documento LIKE ?", "%"+query.Documento+"%").
		Count(&count)

	gs.DB.Model(&guests.Guest{}).
		Where("telefone LIKE ?", "%"+query.Telefone+"%").
		Where("nome LIKE ?", "%"+query.Nome+"%").
		Where("documento LIKE ?", "%"+query.Documento+"%").
		Preload("CheckIns").
		Limit(int(query.Limit)).
		Offset(int((query.Page - 1) * query.Limit)).
		Find(&result.Guests)

	result.TotalPages = (uint(count) / query.Limit)

	if (uint(count) % query.Limit) != 0 {
		result.TotalPages++
	}

	return result
}

// Create and return a GuestRepo
func CreateGuestRepo(db *gorm.DB) guests.GuestRepo {
	return &GuestRepo{
		DB: db,
	}
}

func (gs *GuestRepo) DeleteGuest(guestId uint) error {
	err := gs.DB.Where(&checkin.CheckIn{Hospede: guestId}).Delete(&checkin.CheckIn{Hospede: guestId}).Error
	if err != nil {
		return err
	}
	return gs.DB.Model(&guests.Guest{}).Delete(&guests.Guest{ID: guestId}).Error
}

func (gs *GuestRepo) SearchInHotelGuest(query guests.GuestQuery) struct {
	Guests     []guests.Guest
	TotalPages uint
} {
	// Default Value to Page
	if query.Page == 0 {
		query.Page = 1
	}
	// Default Value to Limit
	if query.Limit == 0 {
		query.Limit = 10
	}

	var result struct {
		Guests     []guests.Guest
		TotalPages uint
	}

	var count int64
	gs.DB.Model(&guests.Guest{}).
		Select("guests.*").
		Joins("JOIN check_ins ON guests.id = check_ins.hospede and check_ins.data_saida is null").
		Where("telefone LIKE ?", "%"+query.Telefone+"%").
		Where("guests.nome LIKE ?", "%"+query.Nome+"%").
		Where("guests.documento LIKE ?", "%"+query.Documento+"%").
		Group("guests.id").
		Count(&count)

	gs.DB.Model(&guests.Guest{}).
		Select("guests.*").
		Joins("JOIN check_ins ON guests.id = check_ins.hospede and check_ins.data_saida is null").
		Where("telefone LIKE ?", "%"+query.Telefone+"%").
		Where("nome LIKE ?", "%"+query.Nome+"%").
		Where("documento LIKE ?", "%"+query.Documento+"%").
		Group("guests.id").
		Preload("CheckIns").
		Limit(int(query.Limit)).
		Offset(int((query.Page - 1) * query.Limit)).
		Find(&result.Guests)

	result.TotalPages = (uint(count) / query.Limit)

	if (uint(count) % query.Limit) != 0 {
		result.TotalPages++
	}

	return result
}
