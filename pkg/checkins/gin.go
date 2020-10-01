package checkins

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func JsonUnmarshalBodyTo(c *gin.Context, a interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(a)
}

// A Gin Handler to Create Checkin
func (cs *CheckInService) CreateCheckInGin(c *gin.Context) {
	var checkin CheckIn
	JsonUnmarshalBodyTo(c, &checkin)
	c.JSONP(200, checkin)
}
