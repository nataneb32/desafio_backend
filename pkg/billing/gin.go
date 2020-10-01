package billing

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
)

func (bs *BillingService) CalculateBillOfGin(c *gin.Context) {
	var request struct {
		CheckInID uint `json:"checkinID"`
	}

	err := gin_helpers.JsonUnmarshalBodyTo(c, &request)

	if err != nil {
		c.AbortWithError(400, c.Error(err))
	}

	err, bill := bs.CalculateBillOf(request.CheckInID)

	if err != nil {
		return
	}

	c.JSON(200, &gin.H{"price": bill})
}
