package guests

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
	"strconv"
)

// A GinHandler to create a new guest
func (gs *GuestService) CreateGuestGin(c *gin.Context) {
	var guest Guest
	err := gin_helpers.JsonUnmarshalBodyTo(c, &guest)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}
	gs.GuestRepo.CreateGuest(&guest)

	c.JSONP(200, &guest)
}

// A GinHandler to delete a guest.
func (gs *GuestService) DeleteGuestGin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	err = gs.GuestRepo.DeleteGuest(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	c.Status(200)
}

// A GinHandler to searh guests.
func (gs *GuestService) SearchGuestGin(c *gin.Context) {
	var query GuestQuery
	err := c.BindQuery(&query)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	result := gs.GuestRepo.SearchGuest(query)
	c.JSON(200, gin.H{"totalPages": result.TotalPages, "guests": result.Guests})
}

//
func (gs *GuestService) SearchGuestInHotelGin(c *gin.Context) {
	var query GuestQuery
	err := c.BindQuery(&query)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	result := gs.GuestRepo.SearchInHotelGuest(query)

	c.JSON(200, gin.H{"guests": result.Guests, "totalPages": result.TotalPages})
}

func (gs *GuestService) SearchGuestOutHotelGin(c *gin.Context) {
	var query GuestQuery
	err := c.BindQuery(&query)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	result := gs.GuestRepo.SearchOutHotelGuest(query)

	c.JSON(200, gin.H{"guests": result.Guests, "totalPages": result.TotalPages})
}

// A GinHandler to get a guest.
func (gs *GuestService) GetGuestGin(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}
	var guest Guest
	guest.ID = uint(userId)

	err = gs.GuestRepo.GetGuest(&guest)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}
	lbill := gs.BillingService.NewestBillOf(guest.CheckIns)
	bill := gs.BillingService.SumBillOf(guest.CheckIns)
	c.JSONP(200, gin.H{"guest": guest, "totalBill": bill, "lastBill": lbill})
}

func (gs *GuestService) UpdateGuestGin(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	var guest Guest
	err = gin_helpers.JsonUnmarshalBodyTo(c, &guest)
	if err != nil {
		return
	}

	err = gs.GuestRepo.UpdateGuest(uint(id), &guest)
	if err != nil {
		return
	}

	c.Status(200)
}
