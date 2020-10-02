package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/billing"
	"nataneb32.live/hospedagem/pkg/checkin"
	"nataneb32.live/hospedagem/pkg/checkins"
	"nataneb32.live/hospedagem/pkg/gorm/checkin"
	"nataneb32.live/hospedagem/pkg/gorm/guest"
	"nataneb32.live/hospedagem/pkg/guests"
)

type App struct {
	DB             *gorm.DB
	GuestService   *guests.GuestService
	GuestRepo      guests.GuestRepo
	CheckInService *checkins.CheckInService
	CheckInRepo    *gorm_checkin.CheckInRepo
	BillingService *billing.BillingService
}

func Start() *App {
	var a App
	a.init_database()
	a.init_repositories()
	a.init_services()
	return &a
}

func (a *App) init_database() {

	dsn := "host=postgres user=test password=test dbname=test_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	a.DB = db

	// Migrate the schema
	db.AutoMigrate(&guests.Guest{})
	db.AutoMigrate(&checkin.CheckIn{})
}

func (a *App) init_repositories() {
	a.CheckInRepo = gorm_checkin.CreateCheckInRepo(a.DB)
	a.GuestRepo = gorm_guest.CreateGuestRepo(a.DB)
}

func (a *App) init_services() {
	// Setting fees to the billing services
	a.BillingService = billing.NewBillingService(1500, 12000, 2000, 15000)
	a.CheckInService = checkins.CreateCheckInService(a.CheckInRepo, a.BillingService)
	a.GuestService = guests.CreateGuestService(a.GuestRepo, a.BillingService)
}
