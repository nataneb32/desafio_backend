package checkins

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
)

// A Gin Handler to Create Checkin
func (cs *CheckInService) CreateCheckInGin(c *gin.Context) {
	var checkin CheckIn
	gin_helpers.JsonUnmarshalBodyTo(c, &checkin)
	c.JSONP(200, checkin)
}
