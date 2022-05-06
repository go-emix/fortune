package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/service/system"
)

func Login(c *gin.Context) {
	var pa system.LoginParam
	err := c.ShouldBindJSON(&pa)
	if err != nil {
		resp.Err(c, common.Resps["fault"].SetError(err))
		return
	}
	succ, err := system.Login(pa)
	if err != nil {
		resp.Err(c, common.Resps["loginFailed"].SetError(err))
		return
	}
	resp.Data(c, succ)
}
