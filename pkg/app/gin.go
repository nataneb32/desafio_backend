package app

import (
	"github.com/gin-gonic/gin"
)

func (a *App) CreateAppHandlersGin() *gin.Engine {
	r := gin.Default()

	r.PUT("/guest/:id", a.GuestService.UpdateGuestGin)
	r.DELETE("/guest/:id", a.GuestService.DeleteGuestGin)
	r.POST("/guest", a.GuestService.CreateGuestGin)
	r.GET("/guests", a.GuestService.SearchGuestGin)
	r.GET("/guests/inhotel", a.GuestService.SearchGuestInHotelGin)
	r.GET("/guests/outhotel", a.GuestService.SearchGuestInHotelGin)
	r.GET("/guest/:userId", a.GuestService.GetGuestGin)
	// r.POST("/guest/:id/checkin", a.CheckInService.DoCheckInGin)
	// r.POST("/checkout", a.CheckInService.DoCheckOutGin)
	r.POST("/checkin", a.CheckInService.CreateCheckInGin)
	r.GET("/checkin/:id/bill", a.CheckInService.CalculateBillGin)
	r.GET("/checkin/:id", a.CheckInService.GetCheckInGin)
	return r
}
