package checkins

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/checkin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
	"strconv"
)

// A Gin Handler to Create Checkin
func (cs *CheckInService) CreateCheckInGin(c *gin.Context) {
	var checkin checkin.CheckIn

	err := gin_helpers.JsonUnmarshalBodyTo(c, &checkin)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	err = cs.CheckInRepo.CreateCheckIn(&checkin)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSONP(200, checkin)
}

// A Gin Handler to Get ChckIn
func (cs *CheckInService) GetCheckInGin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	var checkin checkin.CheckIn
	checkin.ID = uint(id)

	err = cs.CheckInRepo.GetCheckIn(&checkin)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.JSONP(200, checkin)
}

func (cs *CheckInService) DoCheckInGin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	var request struct {
		AdicionalVeiculo bool `json:"adicionalVeiculo"`
	}

	err = gin_helpers.JsonUnmarshalBodyTo(c, &request)
	if err != nil {
		return
	}

	cs.DoCheckIn(uint(id), request.AdicionalVeiculo)
	c.Status(200)
}

func (cs *CheckInService) CalculateBillGin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	err, bill := cs.CalculateBill(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"bill": bill})
}

// Delete Gin Handler
func (cs *CheckInService) DeleteCheckInGin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	err = cs.CheckInRepo.DeleteCheckIn(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}
	c.Status(200)
}
