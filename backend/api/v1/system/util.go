package system

import (
	"github.com/gin-gonic/gin"
)

func getUid(c *gin.Context) int {
	get, _ := c.Get("uid")
	return get.(int)
}
