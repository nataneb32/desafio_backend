package guests

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
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

}

// A GinHandler to searh guests.
func (gs *GuestService) SearchGuestGin(c *gin.Context) {
	var query struct {
		Documento string
		Nome      string
		Limit     uint
		Page      uint
	}
	err := c.BindQuery(&query)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": true, "message": err.Error()})
		return
	}

	result := gs.GuestRepo.SearchGuest(query)

	c.JSONP(200, result)
}

// A GinHandler to get a guest.
func (gs *GuestService) GetGuestGin(c *gin.Context) {

}
