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
