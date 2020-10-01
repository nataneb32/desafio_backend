package checkins

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
)

// A Gin Handler to Create Checkin
func (cs *CheckInService) CreateCheckInGin(c *gin.Context) {
	var checkin CheckIn

	err := gin_helpers.JsonUnmarshalBodyTo(c, &checkin)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	// Since we are creating a checkin, the checkin struct can't have a id.
	checkin.ID = 0

	err = cs.CheckInRepo.CreateCheckIn(&checkin)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSONP(200, checkin)
}

// A Gin Handler to Get ChckIn
func (cs *CheckInService) GetCheckInGin(c *gin.Context) {
	var checkin CheckIn

	err := gin_helpers.JsonUnmarshalBodyTo(c, &checkin)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	err = cs.CheckInRepo.CreateCheckIn
}
