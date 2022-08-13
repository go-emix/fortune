package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/pkg/tianqi"
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

func Tianqi(c *gin.Context) {
	temp, err := tianqi.GetTemp()
	if err != nil {
		resp.Err(c, resp.Resp{
			ErrCode: 2001,
			ErrMsg:  err.Error(),
		})
		return
	}
	resp.Data(c, temp)
}

func I18n(c *gin.Context) {
	en, zh := system.I18n()
	resp.Data(c, map[string]interface{}{
		"zh": zh,
		"en": en,
	})
}
