package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/service/system"
)

func getUid(c *gin.Context) int {
	get, _ := c.Get("uid")
	return get.(int)
}

func getRoles(c *gin.Context) []system.Role {
	get, _ := c.Get("roles")
	return get.([]system.Role)
}

func getRids(roles []system.Role) []int {
	rids := make([]int, 0)
	for _, r := range roles {
		rids = append(rids, r.Id)
	}
	return rids
}

func I18n(c *gin.Context) {
	en, zh := system.I18n()
	resp.Data(c, map[string]interface{}{
		"zh": zh,
		"en": en,
	})
}

func GetUid(c *gin.Context) int {
	get, ok := c.Get("uid")
	if ok {
		return get.(int)
	}
	return 0
}
