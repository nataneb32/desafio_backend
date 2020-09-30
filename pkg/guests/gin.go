package guests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func JsonUnmarshalBodyTo(c *gin.Context, a interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(a)
}

// A GinHandler to create a new guest
func (gs *GuestService) CreateGuestGin(c *gin.Context) {
	var guest Guest
	JsonUnmarshalBodyTo(c, &guest)

	gs.GuestRepo.CreateGuest(&guest)

	c.JSONP(200, &guest)
}
