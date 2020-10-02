package app

import (
	"github.com/gin-gonic/gin"
)

func (a *App) CreateAppHandlersGin() *gin.Engine {
	r := gin.Default()

	r.POST("/guest", a.GuestService.CreateGuestGin)
	r.GET("/guest", a.GuestService.SearchGuestGin)
	r.GET("/guest/:userId", a.GuestService.GetGuestGin)
	// r.POST("/guest/:id/checkin", a.CheckInService.DoCheckInGin)
	// r.POST("/guest/:id/checkout", a.CheckInService.DoCheckOutGin)
	// r.GET("/guest/inhotel", a.GuestService.GetGuestInHotelGin)
	// r.GET("/guest/outhotel", a.GuestService.GetGuestOutHotelGin)
	r.POST("/checkin", a.CheckInService.CreateCheckInGin)

	return r
}
