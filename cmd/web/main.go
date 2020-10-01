package main

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/app"
)

func main() {

	a := app.Start()

	r := gin.Default()
	r.POST("/guests", a.GuestService.CreateGuestGin)
	r.POST("/checkins", a.CheckInService.CreateCheckInGin)
	r.POST("/bill", a.BillingService.CalculateBillOfGin)
	r.POST("checkins/search", a.CheckInService.GetCheckInGin)
	r.Run()
}
