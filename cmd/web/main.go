package main

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/app"
)

func main() {

	a := app.Start()

	r := gin.Default()
	r.POST("/guests", a.GuestService.CreateGuestGin)
	r.Run()
}
