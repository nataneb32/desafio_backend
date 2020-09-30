package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"nataneb32.live/hospedagem/pkg/gorm/guest"
	"nataneb32.live/hospedagem/pkg/guests"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&gorm_guest.GuestSchema{})

	guestRepo := gorm_guest.CreateGuestRepo(db)
	guestService := guests.CreateGuestService(guestRepo)
	r := gin.Default()
	r.POST("/guests", guestService.CreateGuestGin)
	r.Run()
}
