package gin_helpers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func JsonUnmarshalBodyTo(c *gin.Context, a interface{}) error {
	return json.NewDecoder(c.Request.Body).Decode(a)
}
