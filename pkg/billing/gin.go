package billing

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
)

func (bs *BillingService) CalculateBillOfGin(c *gin.Context) {
	var request struct {
		CheckInID uint `json:"checkinID"`
	}

	gin_helpers.JsonUnmarshalBodyTo(c, &request)

	c.JSON(200, &gin.H{"id": request.CheckInID})
}
