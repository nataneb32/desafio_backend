package guests

import (
	"github.com/gin-gonic/gin"
	"nataneb32.live/hospedagem/pkg/gin_helpers"
)

// A GinHandler to create a new guest
func (gs *GuestService) CreateGuestGin(c *gin.Context) {
	var guest Guest
	gin_helpers.JsonUnmarshalBodyTo(c, &guest)

	gs.GuestRepo.CreateGuest(&guest)

	c.JSONP(200, &guest)
}
