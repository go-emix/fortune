package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-emix/fortune/backend/pkg/common"
	"github.com/go-emix/fortune/backend/pkg/i18n"
	"github.com/go-emix/fortune/backend/pkg/jwt"
	"github.com/go-emix/fortune/backend/pkg/resp"
	"github.com/go-emix/fortune/backend/service/system"
)

func LoginInterceptor(c *gin.Context) {
	th := c.GetHeader("token")
	if th == "" {
		resp.Err(c, i18n.NewErr(c, "not_login", nil).Resp())
		return
	}
	token, err := jwt.ParseToken(th)
	if err != nil {
		if err == jwt.TokenExpiredErr {
			resp.Err(c, i18n.NewErr(c, "token_expired", nil).Resp())
			return
		}
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	c.Set("uid", token.UserId)
	roles := make([]int, 0)
	err = common.DB.Model(system.AdminRole{}).Joins("left join admin on "+
		"admin.id=admin_role.admin").Where("admin=?", token.UserId).
		Pluck("role", &roles).Error
	if err != nil {
		resp.Err(c, i18n.NewErr(c, "", err).Resp())
		return
	}
	c.Set("roles", roles)
}

func PermissionInterceptor(c *gin.Context) {
	//uid := getUid(c)
	//system.Menus()
}
