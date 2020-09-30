package checkins

import (
	"github.com/gin-gonic/gin"
)

// A Gin Handler to Create Checkin
func CreateCheckInGin(c *gin.Context) {
	c.JSON(200, &gin.H{"asd": nil})
}
