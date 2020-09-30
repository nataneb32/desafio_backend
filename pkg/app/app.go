package app

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/gorm/guest"
	"nataneb32.live/hospedagem/pkg/guests"
)

type App struct {
	DB           *gorm.DB
	GuestService *guests.GuestService
	GuestRepo    guests.GuestRepo
}

func Start() *App {
	var a App
	a.init_database()
	a.init_repositories()
	a.init_services()
	return &a
}

func (a *App) init_database() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	a.DB = db

	// Migrate the schema
	db.AutoMigrate(&gorm_guest.GuestSchema{})
}

func (a *App) init_repositories() {
	a.GuestRepo = gorm_guest.CreateGuestRepo(a.DB)
}

func (a *App) init_services() {
	a.GuestService = guests.CreateGuestService(a.GuestRepo)
}
